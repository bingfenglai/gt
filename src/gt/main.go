package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "gt/routers"
	"log"

	// 导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	beego.Run()

}

// 初始化方法
func init() {

	connectMysql()

}

// 连接数据库
func connectMysql() {
	mysqlUser, _ := beego.AppConfig.String("mysqlUser")

	mysqlPass, _ := beego.AppConfig.String("mysqlPass")

	mysqlIp, _ := beego.AppConfig.String("mysqlIp")

	mysqlPort, _ := beego.AppConfig.String("mysqlPort")

	mysqlDbname, _ := beego.AppConfig.String("mysqlDbname")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	ds := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlIp + ":" + mysqlPort + ")" + "/" + mysqlDbname + "?charset=utf8&parseTime=true&loc=Local"

	orm.RegisterDataBase("default", "mysql", ds)
	orm.SetMaxIdleConns("default", 20)
	orm.SetMaxOpenConns("default", 20)

	orm.Debug = true

	_, err := orm.NewOrm().Raw("select 1").Exec()

	if err != nil {
		log.Default().Println("执行sql出错", err.Error())
	} else {
		log.Default().Println("数据库连接成功")
	}
}
