package main

import (
	"fmt"
	"github.com/bingfenglai/gt/initialization"
	"github.com/bingfenglai/gt/router"
	// 导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

// 初始化一个http服务对象

func main() {

	router.R.Run(fmt.Sprintf("%s:%d", viper.GetString("server.address"), viper.GetInt("server.port")))

	defer func() {
		initialization.DB.Close()
	}()
}

// 初始化方法
func init() {

	loadConfig()

	initialization.InitDbConfig()

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
