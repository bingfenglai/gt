package storage

import (
	errs "errors"
	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IUserStorage interface {
	SelectOneByUsername(username string) (*entity.User, error)
	SelectOneByEmail(email string) (*entity.User, error)
	Insert(user *entity.User)(uint,error)
}

type userDbStorageImpl struct {
}

func (store *userDbStorageImpl) SelectOneByUsername(username string) (*entity.User, error) {

	user := entity.User{}

	err := global.DB.Where("username = ?", username).First(&user).Error

	if err != nil && err.Error() == gorm.ErrRecordNotFound.Error() {
		zap.L().Info("记录不存在", zap.Error(err))
		return nil, errors.ErrAccountPasswordMismatch
	}

	return &user, err
}

func (store *userDbStorageImpl) SelectOneByEmail(email string) (*entity.User, error) {

	user := entity.User{}
	err := global.DB.Where("email = ?", email).First(&user).Error

	if err == nil || errs.Is(err,gorm.ErrRecordNotFound) {
		return &user,nil
	}

	return &user,err

}


func(store *userDbStorageImpl)Insert(user *entity.User)(uint,error){
	if user==nil {
		return 0,errors.ErrParamsNotNull
	}

	if err := global.DB.Create(user).Error;err!=nil{
		return 0,err
	}

	return user.ID,nil
}
