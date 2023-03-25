package handler

import (
	"github.com/bingfenglai/gt/common/helper"
	"time"

	"github.com/bingfenglai/gt/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// gin默认日志接入zap中间件
func GinZapLogger() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// 请求之前执行
		startTime := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		ctx.Next()
		// 请求之后执行
		cost := time.Since(startTime)
		//ctx.Header("Access-Control-Allow-Origin","*")
		global.Log.Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("ip", ctx.ClientIP()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
		us := helper.ParseUserAgent(ctx.Request.UserAgent())

		global.Log.Info("us", zap.Any("us", us))

	}

}
