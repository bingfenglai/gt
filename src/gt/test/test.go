package test

import (
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/initialization"
	"github.com/bingfenglai/gt/model/cache"
	"github.com/bingfenglai/gt/service"
)

func init() {
	config.LoadConfig()
	initialization.InitAll()
	cache.InitCache()
	service.InitService()
}
