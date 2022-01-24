package initialization

import (
	"encoding/json"
	"log"
	"time"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/conmon/constants"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/model/cache"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)


func InitCacheConfig(){
	b,_ :=json.Marshal(config.Conf.Cache)
	zap.L().Info("缓存配置"+string(b))
	if config.Conf.Cache.CacheType == constants.RedisCache {
		initRedisConfig()
		cache.InitCache()
	}else{
		zap.L().Warn("未配置缓存")
	}
}

func initRedisConfig() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:         config.Conf.Redis.Addr,
		Password:     config.Conf.Redis.Password,
		DB:           config.Conf.Redis.DefaultDb,
		DialTimeout:  time.Duration(config.Conf.Redis.Timeout) * time.Second,
		PoolSize:     config.Conf.Redis.PoolSize,
		MinIdleConns: config.Conf.Redis.MinConn,
		MaxConnAge:   time.Duration(config.Conf.Redis.MaxConn),
		PoolTimeout:  time.Duration(config.Conf.Redis.Timeout) * time.Second,
	})

	ctx := redisClient.Context()

	var count = 0

	for {
		_, err := redisClient.Ping(ctx).Result()

		if err != nil {
			count++
			log.Default().Println(err, config.Conf.Redis)
			if count > 30 {
				panic("redis 初始化失败"+err.Error())
			}
			time.Sleep(1 * 1e9)
			continue

		} else {
			log.Default().Println("redis 连接成功")
			global.RedisClient = redisClient
			break
		}
	}
}
