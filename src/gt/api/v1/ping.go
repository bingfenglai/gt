package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/config"

	"github.com/gin-gonic/gin"
)

// @Summary 健康检查接口
// @Description 健康检查接口
// @Success 200 {string} string  "ok"
// @Router /v1/ping [get]
func Ping(c *gin.Context) {
	p := config.Conf.Server

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"server":  p,
	})
}
