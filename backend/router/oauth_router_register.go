package router

import (
	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/handler"
	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/oauth/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// oauth相关handler注册
func init() {
	registerAuth()
}

func registerAuth() {
	groupOAuth2.Use(handler.GinZapRecovery(true), handler.GinZapLogger(), handler.AuthorizationHandler(R))

	// 认证
	groupOAuth2.Handle(http.MethodPost, "/token", func(ctx *gin.Context) {
		_ = oauth.OAuth2Server.HandleTokenRequest(ctx.Writer, ctx.Request)

		//if err != nil {
		//
		//	ctx.JSON(http.StatusOK, result.Fail(err.Error()))
		//	ctx.Abort()
		//	//return
		//}
	})

	// 授权
	// TODO 先检查当前会话中用户是否完成认证，未完成则先认证，认证后再次重定向到当前接口，再完成授权码获取
	groupOAuth2.Handle(http.MethodPost, "/authorize", func(ctx *gin.Context) {

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

	// 登出
	groupOAuth2.Handle(http.MethodPost, "/logout", func(ctx *gin.Context) {
		user, _ := utils.GetCurrentUser(ctx.Request)

		err := oauth.OAuth2Server.Manager.RemoveAccessToken(ctx, user.AccessToken)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, result.FailWithErr(err))
			return
		}

		ctx.JSON(http.StatusOK, result.Ok("登出成功"))
	})

}
