package v1

import (
	"github.com/bingfenglai/gt/router"
	"net/http"

	"github.com/bingfenglai/gt/conmon/constants"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 链接重定向接口
// @Description 链接重定向接口
// @Success 200 {string} string  "ok"
// @Router /v1/redirection/{code} [get]
func Redirection(ctx *gin.Context){
	code :=ctx.Params.ByName("code")
	zap.L().Info("获取短码："+code)

	if constants.ShortCodeLength!=len(code) {
		ctx.AbortWithStatus(http.StatusNotFound)
		ctx.Writer.WriteString("资源不存在")
		return
	}

	// 301 临时重定向
	ctx.Redirect(http.StatusFound,"https://google.com")

}

func init() {
	router.GetV1().Any("/redirection/:code",Redirection)

}