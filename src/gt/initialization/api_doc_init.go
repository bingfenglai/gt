package initialization

import (
	"bytes"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/docs"
	"github.com/bingfenglai/gt/router"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"os"
	"os/exec"
)

func InitApiConfig() {
	url := ginSwagger.URL(viper.GetString("swagger.url"))

	router.R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func RunSwagCmd() {

	cmd := exec.Command("swag", "init")
	log.Default().Println("Cmd ", cmd.Args)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println(out.String())

	initSwagConfig()

}

func initSwagConfig() {

	docs.SwaggerInfo.Version = config.Conf.Swagger.Version
	docs.SwaggerInfo.Host = config.Conf.Swagger.Host
	//docs.SwaggerInfo.BasePath =config.Conf.Swagger.BasePath
	docs.SwaggerInfo.Schemes = config.Conf.Swagger.Schemes
	docs.SwaggerInfo.Title = config.Conf.Swagger.Title
	docs.SwaggerInfo.Description = config.Conf.Swagger.Description

	log.Default().Println("api 文档初始化成功")
}
