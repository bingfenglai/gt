package cache

import (
	"time"

	"github.com/bingfenglai/gt/common/constants"
	"github.com/bingfenglai/gt/config"
	"go.uber.org/zap"

	"github.com/bingfenglai/gt/global"
)

var redisCacheImpl Cache

var localCacheImpl Cache

// 二级缓存策略 本地加远程redis
var l2CacheImpl Cache

var cacheImpl Cache

type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	SetWithDefaultExpiration(key string, value interface{}) error
	Get(key string, value interface{}) error
	SetWithJson(key string, value interface{}, expiration time.Duration) (bool, string)
	SetWithJsonAndDefaultExpiration(key string, value interface{}) (bool, string)

	GetWithJson(key string) (bool, string)
	Keys(keyPrefix string) (bool, []string)

	Delete(key ...string) (bool, int64)
}

func InitCache() {

	zap.L().Info("初始化cache package")
	if constants.RedisCache == config.Conf.Cache.CacheType {
		redisCacheImpl = newRedisCache(global.RedisClient, time.Minute*time.Duration(config.Conf.Cache.DefaultCacheTime))
		cacheImpl = redisCacheImpl
	} else if constants.LocalCache ==config.Conf.Cache.CacheType {
		localCacheImpl = newLocalCache(time.Duration(config.Conf.Cache.DefaultCacheTime)*time.Minute, 1*time.Minute)
		cacheImpl = localCacheImpl
	}else if constants.L2Cache==config.Conf.Cache.CacheType {
		redisCacheImpl = newRedisCache(global.RedisClient, time.Minute*time.Duration(config.Conf.Cache.DefaultCacheTime))
		localCacheImpl = newLocalCache(time.Duration(config.Conf.Cache.DefaultCacheTime)*time.Minute, 1*time.Minute)
		cacheImpl = newL2Cache(localCacheImpl,redisCacheImpl)
	}else{
		panic("缓存配置不正确：可选值：“local、redis与l2”")
	}

	

}

func GetCacheImpl() Cache {
	return cacheImpl
}
