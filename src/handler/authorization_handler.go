package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/config"

	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/pojo/result"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthorizationHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

		// TODO uri由于路径参数不好匹配，选用HandlerName作为权限标识符
		uri := context.Request.RequestURI

		log.Default().Println(context.HandlerName())

		ok := checkAnonymousUrls(uri)

		if ok {
			context.Next()
			return
		}

		ti, err := oauth.OAuth2Server.Server.ValidationBearerToken(context.Request)

		

		if err != nil {

			// 校验不通过，不再调用后续函数
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.FailWithMsg(err.Error(), "令牌已过期，请重新登录"))
			return
		}
		zap.L().Info("token info", zap.Any("current user", ti.GetUserID()))

		// context.Next()
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
