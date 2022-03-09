package handler

import (
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/oauth2.v3/errors"
)

func EmailVerificationCodeHandler(req *http.Request) (userID string, err error) {

	email := req.FormValue("email")
	code := req.FormValue("code")

	if email == "" || code == "" {
		return "", errors.ErrInvalidRequest
	}
	zap.L().Info("====模拟校验 邮箱验证码=====")

	return "969391", nil

}
