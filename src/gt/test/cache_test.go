package test

import (
	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
	"testing"
)

func TestSet(t *testing.T) {

	ok, s := service.CacheService.SetWithDefaultExpiration("11", 13)

	if !ok {
		t.Log(s)
	}

}

func TestGet(t *testing.T) {
	ok, s := service.CacheService.Get("11")

	if ok {
		zap.L().Info(s)
	}

	zap.L().Info(s, zap.Bool("flag", ok))

}
