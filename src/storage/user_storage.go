package storage

import (
	"github.com/bingfenglai/gt/errors"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/model/entity"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type IUserStorage interface {

	SelectOneByUsername(username string) (*entity.User,error)

}


type UserStorageImpl struct {
	
}

func(store *UserStorageImpl) SelectOneByUsername(username string) (*entity.User,error){

	user := entity.User{}

	err := global.DB.Where("username = ?",username).First(&user).Error

	if err!=nil&&err.Error()==gorm.ErrRecordNotFound.Error(){
		zap.L().Info("记录不存在",zap.Error(err))
		return nil,errors.ErrAccountPasswordMismatch
	}

	return &user,err
}