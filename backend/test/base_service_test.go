package test

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/service"
	"testing"
)

func TestBaseServiceSave(t *testing.T) {

	tenant := entity.CreateTenant(0, "大中华区", "")
	tenant1 := entity.CreateTenant(0, "大中华区1", "")

	err := service.TenantService.GetService().Save(context.Background(), []*entity.Tenant{tenant, tenant1})
	if err != nil {
		t.Error(err)
	}
}
