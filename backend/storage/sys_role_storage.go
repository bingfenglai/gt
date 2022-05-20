package storage

import (
	"errors"

	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"

	"gorm.io/gorm"
)

type roleStorage struct {
}

func (store *roleStorage) GetRolesByUid(uid string) ([]*entity.Role, error) {
	roles := make([]*entity.Role, 0)

	err := global.DB.Raw(`SELECT t2.id,t2.code,t2.name 
	FROM tb_sys_user_role AS t1 LEFT JOIN  tb_sys_role AS t2 ON t1.role_id = t2.id 
	WHERE t1.user_id = ?`, uid).Scan(&roles).Error

	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return roles, nil
	}

	return nil, err
}
