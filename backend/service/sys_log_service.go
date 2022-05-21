package service

import (
	"context"
	"github.com/bingfenglai/gt/domain/params"
)

type sysLogService struct {
	baseService
}

func (svc *sysLogService) Create(ctx context.Context, param *params.SysLogCreateParams) error {
	panic("implement me")
}
