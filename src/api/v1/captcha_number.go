package v1

import "github.com/gin-gonic/gin"

// @Tags 验证码API
// @Summary 获取邮件验证码
// @Description 获取邮件验证码
// @Success 200 {string} string  "ok"
// @Router /v1/captcha/email [get]
func GetEmailCode(ctx *gin.Context) {

}
