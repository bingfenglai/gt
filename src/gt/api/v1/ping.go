package v1

import (
	"net/http"
	"time"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/models/cache"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// @Summary 健康检查接口
// @Description 健康检查接口
// @Success 200 {string} string  "ok"
// @Router /v1/ping [get]
func Ping(c *gin.Context) {
	p := config.Conf.Server

	zap.L().Info("健康检查接口调用")

	cache := cache.NewRedisCache(global.RedisClient,time.Minute*30)

	cache.SetWithDefaultExpiration(c.Request.Host,c.Request.UserAgent())

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"server":  p,
	})
}
