package storage_interface

import "github.com/bingfenglai/gt/domain/entity"

type IUserStorage interface {
	SelectOneByUId(uid int) (*entity.User, error)
	SelectOneByUsername(username string) (*entity.User, error)
	SelectOneByEmail(email string) (*entity.User, error)
	Insert(user *entity.User) (uint, error)
}
