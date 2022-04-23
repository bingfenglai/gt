package server

import (
	"net/http"

	"github.com/go-oauth2/oauth2/v4"
	"go.uber.org/zap"
)

const EmailCode oauth2.GrantType = "email_code"

var customGrantTypes []oauth2.GrantType


type(

	EmailVerificationCodeHandler func(req *http.Request) (userID string, err error)
) 




func init(){

	zap.L().Info("初始化oauth grant_type")
	customGrantTypes = make([]oauth2.GrantType, 0)
	customGrantTypes = append(customGrantTypes, EmailCode)

}