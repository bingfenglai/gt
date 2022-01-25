package cache

import (
	"context"
	"encoding/json"
	"github.com/bingfenglai/gt/conmon/helper"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type redisCache struct {
	redisClient       *redis.Client
	ctx               context.Context
	defaultExpiration time.Duration
}

func newRedisCache(redisClient *redis.Client, defaultExpiration time.Duration) *redisCache {

	return &redisCache{
		redisClient:       redisClient,
		ctx:               redisClient.Context(),
		defaultExpiration: defaultExpiration,
	}
}

func (receiver *redisCache) Set(key string, value interface{}, expiration time.Duration) (bool, string) {

	value, _ = json.Marshal(value)

	s, err := receiver.redisClient.Set(receiver.ctx, key, value, expiration).Result()
	if err != nil {
		zap.L().Error(err.Error())
		return false, err.Error()
	}

	zap.L().Info("redis set value result: " + s)

	return true, ""
}

func (receiver *redisCache) SetWithDefaultExpiration(key string, value interface{}) (bool, string) {
	return receiver.Set(key, value, receiver.defaultExpiration)
}

func (receiver *redisCache) Get(key string) (bool, string) {

	result, err := receiver.redisClient.Get(receiver.ctx, key).Result()

	ok, _ := helper.CheckErr(err)

	return ok, result

}

func (receiver *redisCache) Keys(keyPrefix string) (bool, []string) {

	result, err := receiver.redisClient.Keys(receiver.ctx, keyPrefix).Result()

	ok, _ := helper.CheckErr(err)

	return ok, result

}

func (receiver *redisCache) Delete(key ...string) (bool, int64) {
	result, err := receiver.redisClient.Del(receiver.ctx, key...).Result()

	ok, _ := helper.CheckErr(err)

	return ok, result

}
