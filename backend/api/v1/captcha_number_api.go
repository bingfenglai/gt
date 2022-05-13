package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/common/constants"
	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/router"
	"github.com/bingfenglai/gt/service"
	"github.com/gin-gonic/gin"
)

// @Tags 验证码API
// @Summary 获取邮件验证码
// @Description 获取邮件验证码
// @Success 200 {object} result.Result
// @Router /v1/captcha/email [get]
func GetEmailCode(ctx *gin.Context) {

	email := ctx.Request.FormValue("email")

	if email == "" || helper.VerifyEmailFormat(email) != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail("邮箱格式错误"))
		return
	}

	id, err := service.CaptchaService.GetNumberCode(email, constants.SEND_CODE_EMIAL)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.FailWithErr(err))
		return
	}

	ctx.JSON(http.StatusOK, result.Ok(id))
}

func init() {
	router.GetV1().GET("/captcha/email", GetEmailCode)

}
