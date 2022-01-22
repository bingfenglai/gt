package initialization

import (
	

	"github.com/bingfenglai/gt/router"
	// "github.com/go-openapi/swag"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/bingfenglai/gt/docs"
	"bytes"
	"log"
	"os/exec"
	"os"
)

func InitApiConfig(){
	url := ginSwagger.URL("http://localhost:9527/swagger/doc.json")
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

