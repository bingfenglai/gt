package test

import (
	"context"
	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/service"
	"testing"
)

func TestCreateTenant(t *testing.T) {
	createParams := params.TenantCreateParams{
		Name:   "省厅",
		Remark: "省厅",
	}
	if err := service.TenantService.Create(createParams, context.Background()); err != nil {
		t.Error(err)
	}
}
