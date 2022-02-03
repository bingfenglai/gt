package service

import (
	"errors"
	"github.com/bingfenglai/gt/model/entity"
	"github.com/bingfenglai/gt/pojo/dto"
	"github.com/jinzhu/gorm"
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
	user := entity.User{}
	err := user.FindByUsername(username)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名或密码错误")
	}

	userDto := dto.UserDTO{
		Username: user.Username,
		Password: user.Password,
	}

	return &userDto, nil
}
