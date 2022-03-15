package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var handlerPathMap = make(map[string]string)

var lock sync.RWMutex

func AuthorizationHandler(engine *gin.Engine) gin.HandlerFunc {

	return func(context *gin.Context) {

		initHandlerNamePathMap(engine)

		req := context.Request

		//uri := req.URL.Path

		uri := handlerPathMap[context.HandlerName()]

		if uri == "" {
			uri = req.URL.Path
		}

		log.Default().Println("当前：", "\n方法", context.Request.Method,
			"\n请求路径:", req.URL.Path,
			"\n处理器名称：", context.HandlerName())

		// log.Default().Println("当前请求路径：", context.Request.URL.Path, context.Request.URL.RawPath)

		ok := checkAnonymousUrls(uri)

		if ok {
			// context.Next()
			return
		}

		ti, err := oauth.OAuth2Server.ValidationBearerToken(context.Request)

		if err != nil {

			// 校验不通过，不再调用后续函数
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.FailWithMsg(err.Error(), "令牌已过期，请重新登录"))
			return
		}
		zap.L().Info("token info", zap.Any("current user", ti.GetUserID()))

		us, err := service.UserSessionService.GetSession(ti.GetUserID())

		if err != nil {

			// 校验不通过，不再调用后续函数
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.FailWithMsg(err.Error(), "令牌已过期，请重新登录"))
			return
		}
		flag := false
		for _, api := range us.Apis {
			if api.Method == "*" && api.Uri == "*" {
				flag = true
				break
			}
			if api.Method == req.Method && api.Uri == uri {
				flag = true
				break
			}
		}

		if !flag {
			context.Abort()
			dataMap := make(map[string]string)
			dataMap["uri"] = req.RequestURI
			context.JSON(http.StatusForbidden, result.FailWithMsg("当前账号没有权限访问此接口", dataMap))
		}

		// 校验当前用户是否有权限访问当前接口
		// TODO 基于字典树优化匹配查询

	}
}

func checkAnonymousUrls(uri string) bool {

	zap.L().Info("当前uri", zap.String("uri", uri))
	_, ok := helper.Find(config.Conf.Auth.AnonymousUrls, uri)

	return ok
}

func initHandlerNamePathMap(engine *gin.Engine) {
	lock.Lock()
	if len(handlerPathMap) == 0 {
		zap.L().Info("初始化HandlerNamePathMap")
		routes := engine.Routes()
		for i := 0; i < len(routes); i++ {
			info := routes[i]
			// zap.L().Info(strconv.Itoa(i), zap.Any("method", info.Method), zap.Any("path", info.Path), zap.Any("handler", info.Handler))
			handlerPathMap[info.Handler] = info.Path
		}
	}

	lock.Unlock()

}
