package storage

import (
	"errors"

	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
	"gorm.io/gorm"
)

type apiStorage struct {
}

func (store *apiStorage) GetApisByRoleIds(roleIds []int) ([]*entity.Api, error) {
	apis := make([]*entity.Api, 0)

	if roleIds == nil || len(roleIds) == 0 {
		return apis, nil
	}

	err := global.DB.Raw(`SELECT t2.id,t2.method,t2.uri 
	FROM tb_sys_role_api AS t1 LEFT JOIN tb_sys_api AS t2 ON t1.api_id = t2.id 
	WHERE t1.role_id IN (?) `, roleIds).Scan(&apis).Error

	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return apis, nil
	}

	return nil, err
}
