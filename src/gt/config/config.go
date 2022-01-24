package config

import "log"

var Conf Config

type Config struct {
	Redis    RedisConfig
	Server   ServerConfig
	DataBase DataBaseConfig
	Swagger  SwaggerConfig
	Log      LogConfig
	Cache CacheConfig
}

func init() {
	log.Default().Println("装载配置文件信息")
}
