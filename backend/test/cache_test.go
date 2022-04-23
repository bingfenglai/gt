package test

import (
	"testing"
	"time"

	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
)

func TestSetWithJson(t *testing.T) {

	ok, s := service.CacheService.SetWithJson("11", "Hello World!", time.Minute*18)

	if !ok {
		t.Log(s)
	}

}

func TestGetWithJson(t *testing.T) {
	ok, s := service.CacheService.GetWithJson("11")

	if ok {
		zap.L().Info(s)
	}

	zap.L().Info(s, zap.Bool("flag", ok))

}

func TestGet(t *testing.T) {

	s :=""
	if err := service.CacheService.Get("11", &s); err != nil {
		t.Error(err)
	} else {
		zap.L().Info("成功获取值", zap.Any("", s))
	}

}

