package initialization

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	_ "github.com/bingfenglai/gt/docs"
	"github.com/bingfenglai/gt/router"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

}
