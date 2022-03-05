package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/pojo/result"
	"github.com/bingfenglai/gt/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ThreadCallback(ctx *gin.Context){

	code := ctx.Request.FormValue("code")
	zap.L().Info("code",zap.Any("code ",code))
	dataMap := make(map[string]interface{})
	dataMap["code"] = code
	ctx.JSON(http.StatusOK,result.Ok(dataMap))


}

func init(){

	router.GetV1().Any("thread_callback",ThreadCallback)
}