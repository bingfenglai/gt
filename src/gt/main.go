package main

import (
	"fmt"
	"log"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/conmon/constants"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/initialization"
	"github.com/bingfenglai/gt/router"

	// 导入mysql驱动
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// 初始化一个http服务对象

func main() {

	router.R.Run(fmt.Sprintf("%s:%d", viper.GetString("server.address"), viper.GetInt("server.port")))

	defer func() {
		global.DB.Close()
		global.RedisClient.Close()
	}()
}

// 初始化方法
func init() {

	loadConfig()
	config.Init()

	go func() {
		initialization.InitDbConfig()	
	}()

	go func() {
		initialization.InitRedisConfig()
	}()

	
	go func() {
		log.Default().Printf("active: %s",config.ServerConfigInfo.ActiveProfiles)
		if config.ServerConfigInfo.ActiveProfiles == constants.Dev {
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
}
