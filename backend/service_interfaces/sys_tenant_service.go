package service_interfaces

import (
	"context"
	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/domain/response"
)

type ITenantService interface {
	GetService() Service
	Create(param params.TenantCreateParams, ctx context.Context) error
	List() ([]*response.TenantResponse, error)
}
