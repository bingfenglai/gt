package api

import (
	"net/http"
	"strconv"

	"github.com/bingfenglai/gt/pojo/result"
	"github.com/bingfenglai/gt/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PrintApiAndHandlers(ctx *gin.Context) {

	// for i, v := range ctx.HandlerNames() {
	// 	zap.L().Info(strconv.Itoa(i),zap.Any("handler",v))
	// }

	routesInfo := router.R.Routes()

	for i := 0; i < len(routesInfo); i++ {
		info := routesInfo[i]
		zap.L().Info(strconv.Itoa(i), zap.Any("method", info.Method), zap.Any("path", info.Path), zap.Any("handler", info.Handler))

	}

	

	ctx.JSON(http.StatusOK, result.Ok(nil))
}

func init() {
	router.R.GET("/api/print", PrintApiAndHandlers)
}
