package initialization

import (
	"github.com/bingfenglai/gt/global"
	"go.uber.org/zap"
	"log"
	"net/http"

	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/oauth/handler"

	"github.com/bingfenglai/gt/oauth/server"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
)

// 初始化oauth相关配置
func initOAuth2Server() {

	// 设置oauth管理器
	oauth.OAuth2Server = server.NewDefaultCustomOAuthServer()

	// 禁止GET请求方式认证
	oauth.OAuth2Server.SetAllowGetAccessRequest(false)

	// 设置获取client属性
	// oauth.OAuth2Server.SetClientInfoHandler(server.ClientFormHandler)
	oauth.OAuth2Server.SetClientInfoHandler(handler.ClientInfoHandler)

	// 设置password模式认证处理器
	oauth.OAuth2Server.SetPasswordAuthorizationHandler(handler.PasswordAuthorizationHandler)

	// 设置邮箱验证码认证方式
	oauth.OAuth2Server.SetEmailVerificationCodeHandler(handler.EmailVerificationCodeHandler)

	//设置用户授权处理器
	oauth.OAuth2Server.SetUserAuthorizationHandler(handler.UserAuthorizationHandler)

	// SetClientAuthorizedHandler check the client allows to use this authorization grant type
	oauth.OAuth2Server.SetClientAuthorizedHandler(handler.ClientAuthorizedHandler)

	// 配置错误处理
	oauth.OAuth2Server.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		global.Log.Error("Internal Error:", zap.Any("err", err))
		//log.Println("Internal Error:", err.Error())
		re = errors.NewResponse(errors.ErrInvalidRequest, http.StatusOK)
		re.Description = err.Error()
		re.ErrorCode = 1
		return
	})

	oauth.OAuth2Server.SetResponseErrorHandler(func(re *errors.Response) {
		log.Default().Println("Response Error:", re.Error.Error(), re.ErrorCode, re.StatusCode, re.Description)

	})

	// 设置认证成功后响应的扩展字段
	oauth.OAuth2Server.SetExtensionFieldsHandler(func(ti oauth2.TokenInfo) (fieldsValue map[string]interface{}) {

		fieldsValue = make(map[string]interface{})
		fieldsValue["msg"] = "Welcome to gt. Here is a short link one-stop solution."
		return
	})

	oauth.OAuth2Server.SetResponseTokenHandler(handler.ResponseTokenHandler)

}
