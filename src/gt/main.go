package main

import (
	"fmt"
	"log"

	"github.com/bingfenglai/gt/config"


	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/initialization"
	"github.com/bingfenglai/gt/router"
	"github.com/gin-gonic/gin"

	// 导入mysql驱动
	"os"

	_ "github.com/bingfenglai/gt/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// 初始化一个http服务对象

func main() {

	//router.R.Run(fmt.Sprintf("%s:%d", viper.GetString("server.address"), viper.GetInt("server.port")))

	log.Println(config.Conf.Server)
	gin.SetMode(config.Conf.Server.Mode)
	router.R.Run(fmt.Sprintf("%s:%d", config.Conf.Server.Address, config.Conf.Server.Port))

	defer func() {
		global.DB.Close()
		global.RedisClient.Close()
	}()
}

// 初始化方法
func init() {

	loadConfig()

	go func() {
		initialization.InitLogConfig()
	}()

	go func() {
		initialization.InitDbConfig()
	}()

	go func() {
		initialization.InitCacheConfig()
		initialization.InitService()
	}()

	go func() {
		log.Default().Printf("active: %s", config.Conf.Server.Mode)
		if config.Conf.Server.Mode == gin.DebugMode {
			initialization.RunSwagCmd()

			initialization.InitApiConfig()

		}
	}()

}

// 加载配置文件

func loadConfig() {
	// 获取当前工作目录
	workDir, _ := os.Getwd()
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config.Conf)

	if err != nil {
		log.Default().Println("初始化配置信息失败\n", err.Error())
	} else {
		log.Default().Println("config info:\n", config.Conf.Redis, "\n", config.Conf.Server, "\nlog: ", config.Conf.Log)
	}

}
