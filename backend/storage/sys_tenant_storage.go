package storage

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
)

type ITenantStorage interface {
	Insert(tenant *entity.Tenant, ctx context.Context) error

	SelectAll()([]*entity.Tenant,error)
}

type tenantStorate struct {
}

func (store *tenantStorate) Insert(tenant *entity.Tenant, ctx context.Context) error {

	return global.DB.WithContext(ctx).Create(tenant).Error
}

func (store *tenantStorate) SelectAll()(tenants []*entity.Tenant,err error){

	tenants = make([]*entity.Tenant, 0)
	err = global.DB.Select("id","name").Where(" status = 0").Find(&tenants).Error
	
	return
}
