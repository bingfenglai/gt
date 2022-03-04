package service

import (
	"errors"

	"github.com/bingfenglai/gt/pojo/dto"
	"github.com/bingfenglai/gt/storage"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type IUserService interface {
	FindUserByUsername(username string) (*dto.UserDTO, error)
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) FindUserByUsername(username string) (*dto.UserDTO, error) {

	if username == "" {
		return nil, errors.New("用户名不能为空")
	}
	
	user,err := storage.UserStorage.SelectOneByUsername(username)
	
	

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名或密码错误")
	}

	if err!=nil{
		zap.L().Error("err",zap.Any("err:",err.Error()))

	}
	
	userDto := dto.UserDTO{
		Username: user.Username,
		Password: user.Password,
	}

	return &userDto, err
}
