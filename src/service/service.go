package service

import "github.com/bingfenglai/gt/model/cache"

var CacheService cache.Cache

func InitService(){
	CacheService = cache.GetCacheImpl()
}
