package v1

import (
	"github.com/bingfenglai/gt/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ThreadCallback(ctx *gin.Context){
	state := ctx.Request.FormValue("state")
	zap.L().Info("state",zap.Any("state",state))
	zap.L().Info("params",zap.Any("params",ctx.Params))
}

func init(){

	router.GetV1().Any("thread_callback",ThreadCallback)
}