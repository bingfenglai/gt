package interfaces

import "github.com/bingfenglai/gt/domain/dto"

type IApiService interface {
	GetSessionApisByRoleIds(roleIds []int) ([]*dto.ApiSessionDTO, error)
}
