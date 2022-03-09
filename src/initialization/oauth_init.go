package initialization

import (
	"log"
	"net/http"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/oauth/handler"
	"github.com/bingfenglai/gt/oauth/store"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
)

var oauthManager = manage.NewDefaultManager()

// 初始化oauth相关配置
func initOAuth2Server() {
	
	// 指定client 存储策略
	oauthManager.MapClientStorage(&store.ClientDbStore{})

	// 指定token存储策略
	initOAuthTokenStore()


	// 设置oauth管理器
	oauth.OAuth2Server = server.NewDefaultServer(oauthManager)

	// 禁止GET请求方式认证
	oauth.OAuth2Server.SetAllowGetAccessRequest(false)

	// 设置获取client属性
	// oauth.OAuth2Server.SetClientInfoHandler(server.ClientFormHandler)
	oauth.OAuth2Server.SetClientInfoHandler(handler.ClientInfoHandler)

	// 设置password模式认证处理器
	oauth.OAuth2Server.SetPasswordAuthorizationHandler(handler.PasswordAuthorizationHandler)

	//设置用户授权处理器
	oauth.OAuth2Server.SetUserAuthorizationHandler(handler.UserAuthorizationHandler)


	// SetClientAuthorizedHandler check the client allows to use this authorization grant type
	oauth.OAuth2Server.SetClientAuthorizedHandler(handler.ClientAuthorizedHandler)

	
	// 配置错误处理
	oauth.OAuth2Server.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		re = errors.NewResponse(errors.ErrInvalidRequest,http.StatusOK)
		re.Description = err.Error()
		re.ErrorCode = 1
		// re.Header.Add("Content-Type","application/json; charset=utf-8")
		return
	})

	oauth.OAuth2Server.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error(),re.ErrorCode,re.StatusCode,re.Description)

	})

	// 设置认证成功后响应的扩展字段
	// oauth.OAuth2Server.SetExtensionFieldsHandler( func(ti oauth2.TokenInfo) (fieldsValue map[string]interface{}) {
	// 	fieldsValue = make(map[string]interface{})
	// 	fieldsValue["error_code"] = constants.Normal_Status
	// 	return
	// })

	oauth.OAuth2Server.SetResponseTokenHandler(handler.ResponseTokenHandler)

	

}

func initOAuthTokenStore() {
	// use redis token store
	oauthManager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		DB:       15,
		Password: config.Conf.Redis.Password,
	}))
}
