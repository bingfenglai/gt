package test

import (
	"testing"

	"github.com/bingfenglai/gt/common/helper"
	"go.uber.org/zap"
)

func TestIpHelper(t *testing.T) {

	ip := helper.Ip2Str(3232236294)
	zap.L().Info("ip", zap.Any("", ip))
}
