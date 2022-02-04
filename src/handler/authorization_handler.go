package handler

import (
	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizationHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

		_, err := oauth.OAuth2Server.ValidationBearerToken(context.Request)

		if err != nil {

			// 校验不通过，不再调用后续函数
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.FailWithMsg(err.Error(), "令牌已过期，请重新登录 "))
			return
		}
	}
}
