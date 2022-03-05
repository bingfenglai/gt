package router

import (
	"net/http"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/handler"

	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/gin-gonic/gin"
)

//var R = gin.Default()
var R = gin.New()

var groupV1 = R.Group("/v1")

func init() {

	// 认证
	R.Handle(http.MethodPost, "/oauth2/token", func(ctx *gin.Context) {
		err := oauth.OAuth2Server.HandleTokenRequest(ctx.Writer, ctx.Request)

		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusOK, result.Fail(err.Error()))
			return
		}
	})

	// 授权
	// TODO 先检查当前会话中用户是否完成认证，未完成则先认证，认证后再次重定向到当前接口，再完成授权码获取
	R.Handle(http.MethodPost, "/oauth2/authorize", func(ctx *gin.Context) {
		err := oauth.OAuth2Server.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusOK, result.Fail(err.Error()))
			return
		}
	})

	// 授权码
	// R.Handle(http.MethodGet, "/oauth2/authorize", func(ctx *gin.Context) {
	// 	err := oauth.OAuth2Server.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	// 	if err != nil {
	// 		ctx.Abort()
	// 		ctx.JSON(http.StatusOK, result.Fail(err.Error()))
	// 		return
	// 	}
	// })

	// 鉴权
	groupV1.Use(handler.AuthorizationHandler())

	// 处理404
	R.NoRoute( func(c *gin.Context) {
		c.Redirect(http.StatusFound,config.Conf.Server.Url404)
	})

}

func GetV1() *gin.RouterGroup {
	return groupV1
}
