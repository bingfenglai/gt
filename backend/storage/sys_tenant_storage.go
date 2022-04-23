package storage

import (
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
)

type ITenantStorage interface {
	Insert(tenant *entity.Tenant) error
}

type tenantStorate struct {
}

func (store *tenantStorate) Insert(tenant *entity.Tenant) error {

	return global.DB.Create(tenant).Error
}
