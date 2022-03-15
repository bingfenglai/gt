package v1

import (
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/oauth/utils"
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
	_ = ctx.ShouldBindBodyWith(&genParams, binding.JSON)
	zap.L().Info("接收到参数", zap.Reflect("genParams", genParams))

	_,err := utils.GetCurrentUId(ctx.Request)
	var code *entity.ShortCode
	if err!=nil {
		code, err = service.ShortCodeService.CreateShortCode(genParams.OriginalLink, true, false)
	}else{
		code, err = service.ShortCodeService.CreateShortCode(genParams.OriginalLink, false, true)
	}

	if err!=nil {
		ctx.JSON(http.StatusBadRequest,err)
		return
	}

	ctx.JSON(http.StatusOK, result.Ok(code.ShortCode))

}

func init() {
	router.GetV1().POST("/shortCode", GenShortCode)

}
