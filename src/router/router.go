package router

import (
	"github.com/gin-gonic/gin"

)

//var R = gin.Default()
var R = gin.New()

var groupV1 = R.Group("v1")

func init() {

	

	//groupV1.GET("/hello", func(ctx *gin.Context) {
	//
	//	name := ctx.Query("name")
	//	ctx.JSON(200, gin.H{
	//		"message": "hello " + name,
	//	})
	//})
	//
	//
	//groupV1.Any("/redirection/:code",v1.Redirection)

	// 设置一个get请求的路由，url为/ping, 处理函数（或者叫控制器函数）是一个闭包函数。
	// groupV1.GET("/ping", func(c *gin.Context) {
	// 	// 通过请求上下文对象Context, 直接往客户端返回一个json
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	//groupV1.GET("/ping", v1.Ping)
	//
	//groupV1.POST("/short_code/gen",v1.GenShortCode)
}


func GetV1()(*gin.RouterGroup){
	return groupV1
}
