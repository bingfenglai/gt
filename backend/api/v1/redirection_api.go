package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/router"
	"github.com/bingfenglai/gt/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
// @Tags 短链接API
// @Summary 链接重定向接口
// @Description 链接重定向接口
// @Success 200 {object} result.Result
// @Router /v1/redirection [get]
func Redirection(ctx *gin.Context) {
	code := ctx.Params.ByName("code")
	zap.L().Info("获取短码：" + code)

	if sc, err := service.ShortCodeService.FindLinkByCode(code); err == nil && sc != nil {
		// 301 临时重定向
		ctx.Redirect(http.StatusFound, sc.Original)
		go service.ShortCodeLogService.Create(uint64(sc.ID), ctx.Request.UserAgent(), helper.ClientIP(ctx.Request))

	} else {

		ctx.Redirect(http.StatusFound, config.Conf.Server.Url404)
	}

}

func init() {
	router.GetV1().Any("/redirection/:code", Redirection)

}
