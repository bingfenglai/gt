package config

import (
	

	"github.com/spf13/viper"
)

func Init() {
	
	ServerConfigInfo = ServerConfig{
		Address: viper.GetString("server.address"),
		Port: viper.GetInt("server.port"),
		ActiveProfiles: viper.GetString("server.profiles.active"),

	}
	// port
	RedisConfigInfo = &RedisConfig{
		Addr:      viper.GetString("redis.addr"),
		Password:  viper.GetString("redis.password"),
		DefaultDb: viper.GetInt("redis.defaultDb"),
		Timeout:   viper.GetInt("redis.timeout"),
		PoolSize:  viper.GetInt("redis.poolSiez"),
		MinConn:   viper.GetInt("redis.minConn"),
		MaxConn:   viper.GetInt("redis.maxConn"),
	}

}


func init(){
	println("config init")
}