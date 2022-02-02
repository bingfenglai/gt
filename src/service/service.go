package service

import "github.com/bingfenglai/gt/model/cache"

var CacheService cache.Cache

var CaptchaService ICaptchaService

func InitService() {
	CacheService = cache.GetCacheImpl()
	CaptchaService = &CaptchaServiceImpl{}

}
