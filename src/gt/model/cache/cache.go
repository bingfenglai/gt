package cache

import (
	
	"time"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/conmon/constants"
	"go.uber.org/zap"

	"github.com/bingfenglai/gt/global"
)


var cacheImpl Cache



type Cache interface {
	Set(key string, value interface{},expiration time.Duration) bool
	SetWithDefaultExpiration(key string, value interface{}) bool

	// Get(key string) (bool,interface{})
	// Keys(keyPrefix string) (bool,[]string)

	// Delete(key string)bool
}



func InitCache(){
	
	zap.L().Info("初始化cache package")
	if constants.RedisCache==config.Conf.Cache.CacheType {
		cacheImpl = newRedisCache(global.RedisClient,time.Second*time.Duration(config.Conf.Cache.DefaultCacheTime))
	}else{
		 zap.L().Warn("未配置缓存")
		 
	}

}


func GetCacheImpl() Cache{
	
	return cacheImpl
}



