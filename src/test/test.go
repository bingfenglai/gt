package test

import (
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/initialization"
	"github.com/bingfenglai/gt/model/cache"
)

func init() {
	config.LoadConfig()
	initialization.InitAll()
	cache.InitCache()

}
