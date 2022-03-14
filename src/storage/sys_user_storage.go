package storage

import (
	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IUserStorage interface {
	SelectOneByUsername(username string) (*entity.User, error)
}

type UserStorageImpl struct {
}

func (store *UserStorageImpl) SelectOneByUsername(username string) (*entity.User, error) {

	user := entity.User{}

	err := global.DB.Where("username = ?", username).First(&user).Error

	if err != nil && err.Error() == gorm.ErrRecordNotFound.Error() {
		zap.L().Info("记录不存在", zap.Error(err))
		return nil, errors.ErrAccountPasswordMismatch
	}

	return &user, err
}
