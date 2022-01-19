package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/bingfenglai/gt/controllers"
)

// go 包初始化函数，go语言中在导入一个包的时候，如果被导入包存在init函数，会执行init函数
// 因此这里可以使用init函数初始化路由设置
func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Get("/goodbye", controllers.Goodbye)
}
