package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/router"
	"github.com/gin-gonic/gin"
)

func ThreadCallback(ctx *gin.Context) {

	// code := ctx.Request.FormValue("code")
	// zap.L().Info("code",zap.Any("code ",code))
	// dataMap := make(map[string]interface{})
	// ctx.ShouldBindQuery(&dataMap)
	// ctx.ShouldBindJSON(&dataMap)
	// dataMap["code"] = code
	ctx.JSON(http.StatusOK, result.Ok(ctx.Request.URL.Query()))

	// zap.L().Info("ctx",zap.Any("",))

}

func init() {

	router.GetV1().Any("thread_callback", ThreadCallback)
}
