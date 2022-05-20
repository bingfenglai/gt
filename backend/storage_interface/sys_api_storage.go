package storage_interface

import "github.com/bingfenglai/gt/domain/entity"

type IApiStorage interface {
	GetApisByRoleIds(roleIds []int) ([]*entity.Api, error)
}
