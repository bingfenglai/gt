package test

import (
	"github.com/bingfenglai/gt/domain/entity"
	"testing"

	"github.com/bingfenglai/gt/common/model/session"
	"github.com/bingfenglai/gt/oauth/utils"
	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
)

func TestCreateTenant(t *testing.T) {
	//createParams := params.TenantCreateParams{
	//	Name:   "省厅",
	//	Remark: "省厅",
	//}

	ctx := utils.GtContext{UserSession: &session.UserSessionInfo{Uid: "13", TenantId: "13"}}

	tenant := entity.CreateTenant(0, "大中华区", "")

	if err := service.TenantService.Save(&ctx, tenant); err != nil {
		t.Error(err)
	}
}

func TestList(t *testing.T) {

	//service.TenantService.Save()

	if list, err := service.TenantService.List(); err == nil {
		for _, v := range list {
			zap.L().Info("租户信息\n", zap.Any("tenant", v))
		}
	} else {
		t.Error(err)
	}
}
