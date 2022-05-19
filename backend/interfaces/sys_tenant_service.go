package interfaces

import (
	"context"
	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/domain/response"
)

type ITenantService interface {
	Create(param params.TenantCreateParams, ctx context.Context) error
	List() ([]*response.TenantResponse, error)
}
