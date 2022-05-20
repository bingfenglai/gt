package storage_interface

import "github.com/bingfenglai/gt/domain/entity"

type IRoleStorage interface {
	GetRolesByUid(uid string) ([]*entity.Role, error)
}
