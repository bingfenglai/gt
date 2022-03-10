package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/bingfenglai/gt/router"
	"github.com/bingfenglai/gt/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary 获取图片人机验证码
// @Description 获取图片人机验证码
// @Success 200 {string} string  "ok"
// @Router /v1/captcha/behavioral/images [get]
func GetImagesBehavioralCaptcha(ctx *gin.Context) {

	resp, err := service.CaptchaService.GetImagesBehavioralCaptcha()

	ok, s := helper.CheckErr(err)
	if ok {
		ctx.JSON(http.StatusOK, result.Ok(resp))
		return
	}

	ctx.JSON(http.StatusServiceUnavailable, result.Fail(s))
}

// @Summary 校验图片人机验证码
// @Description 校验图片人机验证码
// @Success 200 {string} string  "ok"
// @Router /v1/captcha/behavioral/images/verity [post]
func Verity(ctx *gin.Context) {
	p := params.VerityCaptchaParams{}
	err := ctx.ShouldBindBodyWith(&p, binding.JSON)

	ok, s := helper.CheckErr(err)

	if !ok {
		ctx.JSON(http.StatusOK, result.Fail(s))
		return
	}

	ok, err = service.CaptchaService.ImagesBehavioralVerify(p.Dots, p.CaptchaId)

	if err == nil {
		ctx.JSON(http.StatusOK, result.Ok(ok))
		return
	}

	ctx.JSON(http.StatusOK, result.Fail(err.Error()))
	
}

func init() {
	router.GetV1().GET("/captcha/behavioral/images", GetImagesBehavioralCaptcha)
	router.GetV1().POST("/captcha/behavioral/images/verity", Verity)
}
