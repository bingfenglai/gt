package handler

import (
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/conmon/helper"
	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func AuthorizationHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

		uri := context.Request.RequestURI

		ok := checkAnonymousUrls(uri)

		if ok {
			context.Next()
			return
		}

		_, err := oauth.OAuth2Server.ValidationBearerToken(context.Request)

		if err != nil {

			// 校验不通过，不再调用后续函数
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.FailWithMsg(err.Error(), "令牌已过期，请重新登录 "))
			return
		}
	}
}

func checkAnonymousUrls(uri string) bool {

	split := strings.Split(uri, "?")

	if len(split) >= 1 {
		uri = split[0]
	}

	zap.L().Info("当前uri", zap.String("uri", uri))
	_, ok := helper.Find(config.Conf.Auth.AnonymousUrls, uri)

	return ok
}
