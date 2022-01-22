package initialization

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	
	"github.com/bingfenglai/gt/docs"
	
	"github.com/bingfenglai/gt/router"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitApiConfig(){
	url := ginSwagger.URL(viper.GetString("swagger.url"))
	
	router.R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func RunSwagCmd(){
	
	cmd := exec.Command("swag", "init")
	log.Default().Println("Cmd ",cmd.Args)
	
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println(out.String())

}

func InitSwagConfig(){
	docs.SwaggerInfo.Version = viper.GetString("swagger.version")
	docs.SwaggerInfo.Host	= viper.GetString("swagger.host")
	docs.SwaggerInfo.BasePath =	viper.GetString("sawgger.basePath")
	docs.SwaggerInfo.Schemes = 	[]string{}
	docs.SwaggerInfo.Title = 	viper.GetString("swagger.title")
	docs.SwaggerInfo.Description = 	viper.GetString("swagger.description")
	log.Default().Fatalln("api 文档初始化成功")
}

