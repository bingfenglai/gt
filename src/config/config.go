package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Conf Config

type Config struct {
	Redis    RedisConfig
	Server   ServerConfig
	DataBase DataBaseConfig
	Swagger  SwaggerConfig
	Log      LogConfig
	Cache    CacheConfig
}

func init() {
	log.Default().Println("装载配置文件信息")
	LoadConfig()
	//zap.L().Info("装载配置文件信息")

}

// 加载配置文件

func LoadConfig() {
	// 获取当前工作目录
	workDir, _ := os.Getwd()
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)

	if err != nil {
		log.Default().Println("初始化配置信息失败\n", err.Error())
	} else {
		log.Default().Println("config info:\n", Conf.Redis, "\n", Conf.Server, "\nlog: ", Conf.Log)
	}

}
