package main

import (
	"fmt"
	"log"

	"github.com/bingfenglai/gt/config"

	"github.com/bingfenglai/gt/global"

	"github.com/bingfenglai/gt/initialization"
	"github.com/bingfenglai/gt/router"

	"github.com/gin-gonic/gin"

	_ "github.com/bingfenglai/gt/service"
	_ "github.com/bingfenglai/gt/api"
	// 初始化api
	_ "github.com/bingfenglai/gt/api/v1"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化一个http服务对象

func main() {

	//router.R.Run(fmt.Sprintf("%s:%d", viper.GetString("server.address"), viper.GetInt("server.port")))

	log.Println(config.Conf.Server)
	gin.SetMode(config.Conf.Server.Mode)
	router.R.Run(fmt.Sprintf("%s:%d", config.Conf.Server.Address, config.Conf.Server.Port))

	defer func() {

		global.RedisClient.Close()
	}()
}

// 初始化方法
func init() {

	initialization.InitAll()

	go func() {
		log.Default().Printf("active: %s", config.Conf.Server.Mode)
		if config.Conf.Server.Mode == gin.DebugMode {
			initialization.RunSwagCmd()

			initialization.InitApiConfig()

		}
	}()

}
