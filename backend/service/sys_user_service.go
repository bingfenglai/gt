package service

import (
	"context"
	"errors"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/global"
	"strconv"
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

	FindUserByUId(uid int) (*dto.UserDTO, error)

	// FindUserByUIdWithCache 先从缓存当中查找，找不到再走数据库并回写缓存
	FindUserByUIdWithCache(uid int) (*dto.UserDTO, error)

	FindUserByEmail(email string) (*dto.UserDTO, error)

	// FindUserByEmailWithRegister 跟据邮箱查找用户信息，如果不存在则进行注册
	FindUserByEmailWithRegister(email string) (*dto.UserDTO, error)

	// UpdatePwd 更新密码
	UpdatePwd(ctx context.Context, p *params.UpdatePasswordParams, uid int) error
}

type userServiceImpl struct {
}


func (svc *userServiceImpl)FindUserByUsername(username string) (*dto.UserDTO, error){
	if username=="" {
		return nil,errors.New("用户名不能为空")
	}
	var user *entity.User
	var err error
	if helper.VerifyEmailFormat(username)==nil {
		user,err = storage.UserStorage.SelectOneByEmail(username)
	}else{
		user,err =storage.UserStorage.SelectOneByUsername(username)
	}


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

func (u *userServiceImpl) FindUserByUId(uid int) (*dto.UserDTO, error) {

	if uid == 0 {
		return nil, errors.New("用户ID不能为空")
	}

	user, err := storage.UserStorage.SelectOneByUId(uid)

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

func (svc *userServiceImpl) FindUserByUIdWithCache(uid int) (*dto.UserDTO, error) {
	user := dto.UserDTO{}
	if CacheService != nil {
		err := CacheService.Get(strconv.Itoa(uid), &user)
		if err == nil {
			zap.L().Info("user_dto", zap.Any("user", user))
			return &user, nil

		}

	}
	dbUser, err := svc.FindUserByUId(1)

	if CacheService != nil && dbUser != nil {
		go CacheService.Set(strconv.Itoa(uid), dbUser, time.Minute*30)
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

	return udto, err
}

func (svc *userServiceImpl) createByEmail(email string) (*dto.UserDTO, error) {
	s :=strings.Split(email, "@")
	
	username:= s[0]+strings.Split(s[1],".")[0]
	user := entity.User{
		Email:    email,
		Username: username,
		Password: "",
	}

	user.CreatedAt = time.Now()

	uid, err := storage.UserStorage.Insert(&user)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		Uid:      int(uid),
		Username: user.Username,
		Password: user.Password,
		TenantId: user.TenantId,
	}, nil
}

func (svc *userServiceImpl) UpdatePwd(ctx context.Context, p *params.UpdatePasswordParams, uid int) (err error) {
	zap.L().Debug("修改密码")

	if err = p.Check();err!=nil{
		return err
	}
	var op string
	var np string
	if config.Conf.Server.Encrypted {
		key := config.Conf.Encrypt.AesKey
		bk := []byte(key)
		op, err = helper.AesDecryptCFB(p.OldPwd, bk)
		if err != nil {
			zap.L().Error(err.Error())
			return errors.New("请使用密文传输凭证信息")
		}
		np, err = helper.AesDecryptCFB(p.NewPwd, bk)
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}
	}else {
		op = p.OldPwd
		np = p.NewPwd
	}

	user := entity.User{}

	err = global.DB.First(&user).Where("id = ?", uid).Error
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	_, err = PasswordEncodeService.Check(op, user.Password)
	if err != nil {
		return errors.New("密码错误")
	}

	encodedPwd, err := PasswordEncodeService.Encode(np)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	user.Password = encodedPwd
	err = global.DB.WithContext(ctx).Updates(&user).Error

	return

}
