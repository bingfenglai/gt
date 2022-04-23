package v1

import (
	"github.com/bingfenglai/gt/service"
	"net/http"

	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/router"
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
	err := ctx.ShouldBindBodyWith(&genParams, binding.JSON)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	zap.L().Info("接收到参数", zap.Reflect("genParams", genParams))

	sc, err := service.ShortCodeService.CreateShortCodeWithContext(genParams, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, result.Ok(sc.ShortCode))

}

func init() {
	router.GetV1().POST("/shortCode", GenShortCode)

}
