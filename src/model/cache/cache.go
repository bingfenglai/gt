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
	Set(key string, value interface{}, expiration time.Duration)error
	SetWithDefaultExpiration(key string, value interface{}) error
	Get(key string,value interface{})error
	SetWithJson(key string, value interface{}, expiration time.Duration) (bool, string)
	SetWithJsonAndDefaultExpiration(key string, value interface{}) (bool, string)

	GetWithJson(key string) (bool, string)
	Keys(keyPrefix string) (bool, []string)

	Delete(key ...string) (bool, int64)
}

func InitCache() {

	zap.L().Info("初始化cache package")
	if constants.RedisCache == config.Conf.Cache.CacheType {
		cacheImpl = newRedisCache(global.RedisClient, time.Second*time.Duration(config.Conf.Cache.DefaultCacheTime))
	} else {
		zap.L().Warn("未配置缓存")

	}

}

func GetCacheImpl() Cache {

	return cacheImpl
}
