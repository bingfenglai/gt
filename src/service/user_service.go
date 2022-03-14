package service

import (
	"errors"
	"time"

	"github.com/bingfenglai/gt/domain/dto"
	"github.com/bingfenglai/gt/storage"
	"go.uber.org/zap"
)

type IUserService interface {
	FindUserByUsername(username string) (*dto.UserDTO, error)

	// 先从缓存当中查找，找不到再走数据库并回写缓存
	FindUserByUsernameWithCache(username string) (*dto.UserDTO, error)
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
