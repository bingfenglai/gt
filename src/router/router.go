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
