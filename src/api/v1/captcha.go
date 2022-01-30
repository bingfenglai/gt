package v1

import (
	"bytes"
	"encoding/base64"
	"github.com/bingfenglai/gt/conmon/constants"
	"github.com/bingfenglai/gt/pojo/response"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/bingfenglai/gt/router"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/dchest/captcha"

// @Summary 获取验证码
// @Description 获取验证码
// @Success 200 {string} string  "ok"
// @Router /v1/captcha [get]
func GetCaptcha(ctx *gin.Context) {

	id :=captcha.New()

	if id=="" {
		r :=result.Fail("生成验证码失败")
		ctx.JSON(http.StatusInternalServerError,r)
		return
	}

	captchaResponse := response.CaptchaResponse{CaptchaId: id}
	var imagesBuf bytes.Buffer
	_ = captcha.WriteImage(&imagesBuf, id, 800, 400)
	data := base64.StdEncoding.EncodeToString(imagesBuf.Bytes())
	data = constants.ImageBase64Prefix+data


	captchaResponse.ImageUrl = data

	ok := result.Ok(captchaResponse)
	ctx.JSON(http.StatusOK,ok)
}

func init() {
	router.GetV1().GET("/captcha",GetCaptcha)
}
