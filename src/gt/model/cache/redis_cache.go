package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type redisCache struct {
	redisClient *redis.Client
	ctx context.Context
	defaultExpiration time.Duration
	
}

func newRedisCache(redisClient *redis.Client,defaultExpiration time.Duration) *redisCache {

	return &redisCache{
		redisClient: redisClient,
		ctx: redisClient.Context(),
		defaultExpiration: defaultExpiration,
	}
}


func (receiver *redisCache) Set(key string, value interface{},expiration time.Duration) bool{
	
	s,err:= receiver.redisClient.Set(receiver.ctx,key,value,expiration).Result()
	if err!=nil {
		zap.L().Error(err.Error())
		return false
	}

	zap.L().Info("redis set value result: "+s)

	return true
}

func (receiver *redisCache) SetWithDefaultExpiration(key string, value interface{}) bool{
	return receiver.Set(key,value,receiver.defaultExpiration)
}