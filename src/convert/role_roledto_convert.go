package convert

import (
	"github.com/bingfenglai/gt/domain/dto"
	"github.com/bingfenglai/gt/domain/entity"
)

func Role2RoleSessionDTO(role *entity.Role)(*dto.RoleSessionDTO,error){

	if role==nil {
		return nil,nil
	}

	return nil,nil

}

func Role2RoleSessionDTOList(roles []*entity.Role)[]*dto.RoleSessionDTO{
	roledtos := make([]*dto.RoleSessionDTO,len(roles))

	for _, role := range roles {
		
		dto :=dto.RoleSessionDTO{
			RoleId: int(role.ID),
			Code: role.Code,
		}

		roledtos = append(roledtos, &dto)
	}

	return roledtos
}