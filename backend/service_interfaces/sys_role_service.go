package service_interfaces

import "github.com/bingfenglai/gt/domain/dto"

type IRoleService interface {
	GetSessionRolesByUid(uid string) ([]*dto.RoleSessionDTO, error)
}
