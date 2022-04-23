package router

import (
	"net/http"
	"strings"

	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/oauth"
	"github.com/gin-gonic/gin"
)

// oauth相关handler注册
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

		token := ctx.GetHeader("Authorization")

		if token == "" || !strings.HasPrefix(token, "Bearer") {
			// 模拟跳转登录页面

			http.Redirect(ctx.Writer, ctx.Request, "/v1/ping", http.StatusFound)
			return
		}

		err := oauth.OAuth2Server.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusOK, result.Fail(err.Error()))
			return
		}
	})

}
