package config

import "github.com/spf13/viper"

var RedisConfigInfo *RedisConfig

type RedisConfig struct {
	Addr     string
	Password string
	DefaultDb       int
	Timeout int
	PoolSize int
	MinConn int
	MaxConn int

}

func init() {
	
	// port 

	RedisConfigInfo = &RedisConfig{
		Addr: viper.GetString("redis.addr"),	
		Password: viper.GetString("redis.password"),
		DefaultDb: viper.GetInt("redis.defaultDb"),
		Timeout: viper.GetInt("redis.timeout"),
		PoolSize: viper.GetInt("redis.poolSiez"),
		MinConn: viper.GetInt("redis.minConn"),
		MaxConn: viper.GetInt("redis.maxConn"),
		
	}

}
