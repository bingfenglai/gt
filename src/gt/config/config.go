package config

var Conf Config

type Config struct {
	Redis    RedisConfig
	Server   ServerConfig
	DataBase DataBaseConfig
	Swagger  SwaggerConfig
}

func init() {
	println("config init")
}
