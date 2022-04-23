package convert

import (
	"github.com/bingfenglai/gt/domain/dto"
	"github.com/bingfenglai/gt/domain/entity"
)

func Apis2SessionApiDTOList(apis []*entity.Api)[]*dto.ApiSessionDTO{

	apiDtos := make([]*dto.ApiSessionDTO,len(apis))

	for i, api := range apis {
		
		adto := dto.ApiSessionDTO{
			Id: int(api.ID),
			Uri: api.Uri,
			Method: api.Method,
		}

		apiDtos[i] = &adto
	}

	return apiDtos

}