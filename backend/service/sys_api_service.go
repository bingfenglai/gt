package service

import (
	"github.com/bingfenglai/gt/convert"
	"github.com/bingfenglai/gt/domain/dto"
	"github.com/bingfenglai/gt/storage"
)

type apiServiceImpl struct {
}

func (svc *apiServiceImpl) GetSessionApisByRoleIds(roleIds []int) ([]*dto.ApiSessionDTO, error) {
	apiSessionDto := make([]*dto.ApiSessionDTO, 0)

	if roleIds == nil || len(roleIds) == 0 {
		return apiSessionDto, nil

	}

	apis, err := storage.ApiStorage.GetApisByRoleIds(roleIds)

	if err != nil {
		return nil, err
	}

	apiSessionDto = convert.Apis2SessionApiDTOList(apis)

	return apiSessionDto, nil
}
