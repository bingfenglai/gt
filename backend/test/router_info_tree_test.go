package test

import (
	"github.com/bingfenglai/gt/router"
	"go.uber.org/zap"
	"strconv"
	"testing"
)

func TestRouterTree(t *testing.T) {
	routesInfo := router.R.Routes()

	for i := 0; i < len(routesInfo); i++ {
		info := routesInfo[i]

		zap.L().Info(strconv.Itoa(i), zap.Any("method", info.Method), zap.Any("path", info.Path), zap.Any("handler", info.Handler))

	}


}
