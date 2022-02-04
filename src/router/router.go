package router

import (
	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

//var R = gin.Default()
var R = gin.New()

var groupV1 = R.Group("/v1")

func init() {

	// 获取令牌
	R.Handle(http.MethodPost, "/oauth2/token", func(ctx *gin.Context) {
		err := oauth.OAuth2Server.HandleTokenRequest(ctx.Writer, ctx.Request)

		if err != nil {
			ctx.JSON(http.StatusOK, result.Fail(err.Error()))
		}
	})

}

func GetV1() *gin.RouterGroup {
	return groupV1
}
