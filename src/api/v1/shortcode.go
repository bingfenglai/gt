package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/router"

	"github.com/bingfenglai/gt/model/shortcodegen"
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

// @Tags 短链接API
// @Summary 生成短链接
// @Description 生成短链接
// @Success 200 {string} string  "ok"
// @Router /v1/shortCode [post]
func GenShortCode(ctx *gin.Context) {
	genParams := params.GenShortCodeParams{}
	_ = ctx.ShouldBindBodyWith(&genParams, binding.JSON)
	zap.L().Info("接收到参数", zap.Reflect("genParams", genParams))

	gen, _ := shortcodegen.GetShortCodeGeneratorByMethod(shortcodegen.Md5Gen)
	codes, _ := gen.GenShortCode(genParams.OriginalLink)
	ctx.JSON(http.StatusOK, result.Ok(codes))

}

func init() {
	router.GetV1().POST("/shortCode", GenShortCode)

}
