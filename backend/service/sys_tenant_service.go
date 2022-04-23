package service

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/storage"

	"github.com/bingfenglai/gt/domain/params"
)

type ITenantService interface {
	Create(param params.TenantCreateParams, ctx context.Context) error
}

type tenantService struct {
}

func (svc *tenantService) Create(param params.TenantCreateParams, ctx context.Context) (err error) {
	if err = param.Check(); err != nil {
		return err
	}

	tenant := entity.CreateTenant(param.Name, param.Remark)

	return storage.TenantStorage.Insert(tenant, ctx)
}
