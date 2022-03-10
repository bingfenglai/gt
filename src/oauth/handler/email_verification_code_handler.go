package handler

import (
	"net/http"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/errors"
	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
)

func EmailVerificationCodeHandler(req *http.Request) (userID string, err error) {

	email := req.FormValue("email")
	code := req.FormValue("code")
	captchaId := req.FormValue("captcha_id")

	if err = helper.VerifyEmailFormat(email);err!=nil{
		return "",err
	}

	if email == "" || code == "" || captchaId==""{
		return "", errors.ErrEmailCodeInvalid
	}

	zap.L().Info("====校验 邮箱验证码=====")

	if err = service.CaptchaService.NumberCodeVerify(code,captchaId,email);err!=nil{
		return "",err
	}

	return "969391", nil

}
