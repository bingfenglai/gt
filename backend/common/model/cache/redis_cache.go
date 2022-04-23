package cache

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/config"

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

func (receiver *redisCache) Set(key string, value interface{}, expiration time.Duration) error {

	if key == "" {
		return errors.New("缓存key不能为空字符串")
	}
	err := receiver.redisClient.Set(receiver.ctx, key, value, expiration).Err()

	if err != nil {
		zap.L().Error("redis缓存设值错误", zap.Error(err))
	}

	return err

}

func (receiver *redisCache) SetWithDefaultExpiration(key string, value interface{}) error {
	zap.L().Info("redis set key: " + key)
	return receiver.Set(key, value, receiver.defaultExpiration)
}

func (receiver *redisCache) SetWithJson(key string, value interface{}, expiration time.Duration) (bool, string) {

	value, _ = json.Marshal(value)

	s, err := receiver.redisClient.Set(receiver.ctx, key, value, expiration).Result()
	if err != nil {
		zap.L().Error(err.Error())
		return false, err.Error()
	}

	zap.L().Info("redis set value result: " + s)

	return true, ""
}

func (receiver *redisCache) SetWithJsonAndDefaultExpiration(key string, value interface{}) (bool, string) {

	return receiver.SetWithJson(key, value, receiver.defaultExpiration)
}

func (receiver *redisCache) Get(key string, value interface{}) error {
	err := receiver.redisClient.Get(receiver.ctx, key).Scan(value)

	if err != nil {
		zap.L().Error("获取缓存值失败", zap.Error(err))
	}

	return err
}

func (receiver *redisCache) GetWithJson(key string) (bool, string) {

	result, err := receiver.redisClient.Get(receiver.ctx, key).Result()

	//receiver.redisClient.GetWithJson(receiver.ctx,"").Scan()

	ok, _ := helper.CheckErr(err)

	return ok, result

}

func (receiver *redisCache) Keys(keyPrefix string) (bool, []string) {

	result, _ := receiver.redisClient.Keys(receiver.ctx, keyPrefix).Result()

	// ok, _ := helper.CheckErr(err)

	return len(result) > 0, result

}

func (receiver *redisCache) Delete(key ...string) (bool, int64) {
	result, err := receiver.redisClient.Del(receiver.ctx, key...).Result()

	ok, _ := helper.CheckErr(err)

	return ok, result

}

func (receiver *redisCache) GetKeyExpiredEventPubSub() *redis.PubSub{

	redisConfig := config.Conf.Redis
	return receiver.redisClient.Subscribe(receiver.ctx, "__keyevent@"+strconv.Itoa(redisConfig.DefaultDb)+"__:expired")

}

func(receiver *redisCache) GetKeySetEventPubSub()*redis.PubSub{
	redisConfig := config.Conf.Redis
	return receiver.redisClient.PSubscribe(receiver.ctx, "__keyevent@"+strconv.Itoa(redisConfig.DefaultDb)+"__:set")

}

func(receiver *redisCache) GetKeyDelEventPubSub()*redis.PubSub{
	redisConfig := config.Conf.Redis
	return receiver.redisClient.PSubscribe(receiver.ctx, "__keyevent@"+strconv.Itoa(redisConfig.DefaultDb)+"__:del")

}
