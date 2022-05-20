package storage

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
)

type tenantStorage struct {
}

func (store *tenantStorage) Insert(tenant *entity.Tenant, ctx context.Context) error {

	return global.DB.WithContext(ctx).Create(tenant).Error
}

func (store *tenantStorage) SelectAll() (tenants []*entity.Tenant, err error) {

	tenants = make([]*entity.Tenant, 0)
	err = global.DB.Select("id", "name").Where(" status = 0").Find(&tenants).Error

	return
}
