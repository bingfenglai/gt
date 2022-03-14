package test

import (
	"github.com/bingfenglai/gt/common/model/cache"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/initialization"
)

func init() {
	config.LoadConfig()
	initialization.InitAll()
	cache.InitCache()

}
