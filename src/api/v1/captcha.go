package v1

import (
	"github.com/bingfenglai/gt/conmon/helper"
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/bingfenglai/gt/router"
	"github.com/bingfenglai/gt/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// @Summary 获取验证码
// @Description 获取验证码
// @Success 200 {string} string  "ok"
// @Router /v1/captcha [get]
func GetCaptcha(ctx *gin.Context) {

	resp, err := service.CaptchaService.GetImagesBehavioralCaptcha()

	ok, s := helper.CheckErr(err)
	if ok {
		ctx.JSON(http.StatusOK, result.Ok(resp))
		return
	}

	ctx.JSON(http.StatusServiceUnavailable, result.Fail(s))
}

func Verity(ctx *gin.Context) {
	p := params.VerityCaptchaParams{}
	err := ctx.ShouldBindBodyWith(&p, binding.JSON)

	ok, s := helper.CheckErr(err)

	if !ok {
		ctx.JSON(http.StatusBadRequest, result.Fail(s))
		return
	}

	ok, err = service.CaptchaService.Verify(p.Dots, p.CaptchaId)

	if err == nil {
		ctx.JSON(http.StatusOK, result.Ok(ok))
		return
	}

	ctx.JSON(http.StatusOK, result.OkWithMsg(ok, err.Error()))
	return
}

func init() {
	router.GetV1().GET("/captcha", GetCaptcha)
	router.GetV1().POST("/captcha/verity", Verity)
}
