package handler

import (
	"net/http"
	"strconv"

	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/common/helper"

	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
)

func EmailVerificationCodeHandler(req *http.Request) (userID string, err error) {

	email := req.FormValue("email")
	code := req.FormValue("code")
	captchaId := req.FormValue("captcha_id")

	if err = helper.VerifyEmailFormat(email); err != nil {
		return "", err
	}

	if email == "" || code == "" || captchaId == "" {
		return "", errors.ErrEmailCodeInvalid
	}

	zap.L().Info("====校验 邮箱验证码=====")

	if err = service.CaptchaService.NumberCodeVerify(code, captchaId, email); err != nil {
		return "", err
	}

	user,err:=service.UserService.FindUserByEmailWithRegister(email)
	return strconv.Itoa(user.Uid), err

}
