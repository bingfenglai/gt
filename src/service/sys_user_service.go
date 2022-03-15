package service

import (
	"errors"
	"strings"
	"time"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/domain/dto"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/storage"
	"go.uber.org/zap"
)

type IUserService interface {
	FindUserByUsername(username string) (*dto.UserDTO, error)

	// 先从缓存当中查找，找不到再走数据库并回写缓存
	FindUserByUsernameWithCache(username string) (*dto.UserDTO, error)

	FindUserByEmail(email string) (*dto.UserDTO, error)

	// 跟据邮箱查找用户信息，如果不存在则进行注册
	FindUserByEmailWithRegister(email string) (*dto.UserDTO, error)
}

type userServiceImpl struct {
}

func (u *userServiceImpl) FindUserByUsername(username string) (*dto.UserDTO, error) {

	if username == "" {
		return nil, errors.New("用户名不能为空")
	}

	user, err := storage.UserStorage.SelectOneByUsername(username)

	if err != nil {
		zap.L().Error("err", zap.Any("err:", err.Error()))
		return nil, err
	}

	userDto := dto.UserDTO{
		Uid:      int(user.ID),
		TenantId: user.TenantId,
		Username: user.Username,
		Password: user.Password,
	}

	return &userDto, err
}

func (svc *userServiceImpl) FindUserByUsernameWithCache(username string) (*dto.UserDTO, error) {
	user := dto.UserDTO{}
	if CacheService != nil {
		err := CacheService.Get(username, &user)
		if err == nil {
			zap.L().Info("user_dto", zap.Any("user", user))
			return &user, nil

		}

	}
	dbUser, err := svc.FindUserByUsername(username)

	if CacheService != nil && dbUser != nil {
		go CacheService.Set(username, dbUser, time.Minute*30)
	}
	return dbUser, err

}

func (svc *userServiceImpl) FindUserByEmail(email string) (*dto.UserDTO, error) {

	if err := helper.VerifyEmailFormat(email); err != nil {
		return nil, err
	}

	user, err := storage.UserStorage.SelectOneByEmail(email)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		Uid:      int(user.ID),
		TenantId: user.TenantId,
		Username: user.Username,
		Password: "-",
	}, nil

}

func (svc *userServiceImpl) FindUserByEmailWithRegister(email string) (*dto.UserDTO, error) {
	if err := helper.VerifyEmailFormat(email); err != nil {
		return nil, err
	}

	udto, err := svc.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if udto.Uid == 0 && udto.Username == "" {
		return svc.createByEmail(email)
	}

	return udto,err
}


func(svc *userServiceImpl) createByEmail(email string)(*dto.UserDTO,error){

	user :=entity.User{
		Email: email,
		Username: strings.Split(email,"@")[0],
		Password: "",

	}

	user.CreatedAt = time.Now()

	uid,err:= storage.UserStorage.Insert(&user)
	if err!=nil {
		return nil,err
	}

	return &dto.UserDTO{
		Uid: int(uid),
		Username: user.Username,
		Password: user.Password,
		TenantId: user.TenantId,
	},nil
}
