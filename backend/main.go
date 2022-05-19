package main

import (
	"fmt"
	"github.com/bingfenglai/gt/config"
	"log"

	"github.com/bingfenglai/gt/global"

	"github.com/bingfenglai/gt/initialization"
	"github.com/bingfenglai/gt/router"

	"github.com/gin-gonic/gin"

	_ "github.com/bingfenglai/gt/api"
	_ "github.com/bingfenglai/gt/service"

	// 初始化api
	_ "github.com/bingfenglai/gt/api/v1"
	_ "github.com/go-sql-driver/mysql"
)

// @title GT API
// @version 1.0
// @description GT 后端接口文档
// @termsOfService https://github.com/bingfenglai

// @contact.name Ferryman
// @contact.url https://github.com/bingfenglai
// @contact.email bingfenglai.dev@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9527
// @BasePath
func main() {

	//router.R.Run(fmt.Sprintf("%s:%d", viper.GetString("server.address"), viper.GetInt("server.port")))

	//log.Println(config.Conf.Server)
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
	log.Default().Println("Application startup complete! 应用启动完成！")

}
