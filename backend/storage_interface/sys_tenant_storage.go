package storage_interface

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
)

type ITenantStorage interface {
	Insert(tenant *entity.Tenant, ctx context.Context) error

	SelectAll() ([]*entity.Tenant, error)
}
