package test

import (
	"github.com/bingfenglai/gt/common/model/session"
	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/oauth/utils"
	"github.com/bingfenglai/gt/service"
	"testing"
)

func TestCreateTenant(t *testing.T) {
	createParams := params.TenantCreateParams{
		Name:   "省厅",
		Remark: "省厅",
	}

	ctx := utils.GtContext{UserSession: &session.UserSessionInfo{Uid: "13", TenantId: "13"}}
	if err := service.TenantService.Create(createParams, &ctx); err != nil {
		t.Error(err)
	}
}
