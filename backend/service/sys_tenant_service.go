package service

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/domain/response"
	"github.com/bingfenglai/gt/service_interfaces"
	"go.uber.org/zap"

	"github.com/bingfenglai/gt/storage"

	"github.com/bingfenglai/gt/domain/params"
)

type tenantService struct {
	baseService
}

func (svc *tenantService) Create(param params.TenantCreateParams, ctx context.Context) (err error) {
	if err = param.Check(); err != nil {
		return err
	}

	zap.L().Info("创建租户入参", zap.Any("param", param))
	tenant := entity.CreateTenant(param.ParentId, param.Name, param.Remark)

	return storage.TenantStorage.Insert(tenant, ctx)
}

func (svc *tenantService) List() (list []*response.TenantResponse, err error) {

	if tenants, err := storage.TenantStorage.SelectAll(); err == nil {
		list = make([]*response.TenantResponse, len(tenants))
		for i, tenant := range tenants {
			tr := &response.TenantResponse{Id: tenant.ID, Name: tenant.Name}
			list[i] = tr
		}
	}

	return
}

func (svc tenantService) GetService() service_interfaces.Service {
	return &svc.baseService
}
