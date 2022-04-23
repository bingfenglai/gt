package storage

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
)

type ITenantStorage interface {
	Insert(tenant *entity.Tenant, ctx context.Context) error
}

type tenantStorate struct {
}

func (store *tenantStorate) Insert(tenant *entity.Tenant, ctx context.Context) error {

	return global.DB.WithContext(ctx).Create(tenant).Error
}
