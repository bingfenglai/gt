package storage

import (
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/model/entity"
)

type IUserStorage interface {

	SelectOneByUsername(username string) (*entity.User,error)

}


type UserStorageImpl struct {
	
}

func(store *UserStorageImpl) SelectOneByUsername(username string) (*entity.User,error){

	user := entity.User{}

	err := global.DB.Where("username = ?",username).First(&user).Error

	return &user,err
}