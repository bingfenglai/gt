package test

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/storage"
	"go.uber.org/zap"
	"testing"
)

func TestFindOne(t *testing.T) {

	te := &entity.Tenant{}
	ma := make(map[string]interface{}, 0)
	ma["name"] = "亚洲信息中心"

	err := storage.TenantStorage.FindOne(context.Background(), te, ma, []string{"name"})
	if err != nil {
		zap.L().Error("err", zap.Error(err))
	}
	zap.L().Info("租户：", zap.Any("tenant", te))

}

func TestSave(t *testing.T) {
	tns := getTenants()

	err := storage.TenantStorage.Save(context.Background(), tns, tns)
	if err != nil {
		zap.L().Error("err", zap.Error(err))
	}
}

func getTenants() []*entity.Tenant {

	tns := make([]*entity.Tenant, 0)

	t1 := entity.CreateTenant(0, "亚洲信息中心", "")

	tns = append(tns, t1)

	t2 := entity.CreateTenant(0, "欧洲信息中心", "")

	tns = append(tns, t2)

	t3 := entity.CreateTenant(0, "美洲信息中心", "")

	tns = append(tns, t3)

	t4 := entity.CreateTenant(0, "澳洲信息中心", "")

	tns = append(tns, t4)

	return tns
}

func TestDel(t *testing.T) {

	ids := []int{1, 2, 3}
	err := storage.TenantStorage.Delete(context.Background(), &entity.Tenant{}, ids)
	if err != nil {
		t.Error(err)
	}

}
