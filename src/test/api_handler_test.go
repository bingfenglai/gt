package test

import (
	"testing"

	_ "github.com/bingfenglai/gt/api/v1"
	"github.com/bingfenglai/gt/router"
	
	"go.uber.org/zap"
)


func TestSaveApiInfo(t *testing.T){


	for _, handler := range router.R.Handlers {
		zap.L().Info("api",zap.Any("handler",handler))
	}



}