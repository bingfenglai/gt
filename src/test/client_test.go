package test

import (
	"github.com/bingfenglai/gt/model/entity"
	"go.uber.org/zap"
	"testing"
)

func TestAddClient(t *testing.T) {

	client := entity.Client{
		ClientBizId: "000000",
		Domain:      "http://localhost",
		Secret:      "999999",
		CreatedBy:   0,
		UpdatedBy:   0,
		Status:      0,
		Remark:      "默认客户端",
	}

	ok, err := client.Create()

	if !ok {
		zap.L().Error(err.Error())
		t.Fail()
	}

}
