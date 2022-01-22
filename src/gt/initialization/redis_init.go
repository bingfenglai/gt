package initialization

import (
	"log"
	"time"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/global"
	"github.com/go-redis/redis/v8"
)

func InitRedisConfig() {

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisConfigInfo.Addr,
		Password: config.RedisConfigInfo.Password,
		DB: config.RedisConfigInfo.DefaultDb,
		DialTimeout: time.Duration(config.RedisConfigInfo.Timeout),
		PoolSize: config.RedisConfigInfo.PoolSize,
		MinIdleConns: config.RedisConfigInfo.MinConn,
		MaxConnAge: time.Duration(config.RedisConfigInfo.MaxConn),
		PoolTimeout: time.Duration(config.RedisConfigInfo.Timeout),
		
	})
	
	ctx := redisClient.Context()
	
	var count = 0

	for {
		_,err := redisClient.Ping(ctx).Result()

		if err!=nil {
			count++
			if count>3 {
				panic("reids 初始化失败"+err.Error())
			}
			continue
			
		
		}else{
			log.Default().Println("redis 连接成功")
			global.RedisClient = redisClient
			break
		}
	}
}