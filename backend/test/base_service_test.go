package test

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/service"
	"testing"
)

func TestBaseServiceSave(t *testing.T) {

	tenant := entity.CreateTenant(0, "大中华区", "")
	service.TenantService.GetService().Save(context.Background(), tenant)
}
