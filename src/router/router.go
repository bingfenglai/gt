package router

import (
	"net/http"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/handler"

	"github.com/gin-gonic/gin"
)

//var R = gin.Default()
var R = gin.New()

var groupV1 = R.Group("/v1")

func init() {

	// 崩溃处理与访问控制
	R.Use(handler.GinZapRecovery(true), handler.GinZapLogger(), handler.AuthorizationHandler(R))
	
	// 处理404
	R.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, config.Conf.Server.Url404)
	})

	groupV1.Use(handler.GinZapRecovery(true), handler.GinZapLogger(), handler.AuthorizationHandler(R))

	R.StaticFile("/403.html", "./statics/403.html")

	R.Handle(http.MethodGet,"/favicon.ico",func(c *gin.Context) {

		c.Redirect(http.StatusFound,config.Conf.Server.Urlfavicon)
	})
}

func GetV1() *gin.RouterGroup {
	return groupV1
}
