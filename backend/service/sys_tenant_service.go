package service

import (
	"context"

	"github.com/bingfenglai/gt/domain/params"
)

type ITenantService interface {
	Create(param params.TenantCreateParams,ctx context.Context) error
}

type tenantService struct {
	
}

func (svc *tenantService) Create(param params.TenantCreateParams,ctx context.Context) (err error)  {
	return
}