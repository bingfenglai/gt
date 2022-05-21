package service_interfaces

import (
	"context"
	"github.com/bingfenglai/gt/domain/params"
)

type ISysLogService interface {
	Service
	Create(ctx context.Context, params *params.SysLogCreateParams) error
}
