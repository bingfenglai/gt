package initialization

import (
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/oauth/store"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	"log"
)

var oauthManager = manage.NewDefaultManager()

// 初始化oauth相关配置
func initOAuth2Server() {

	//manager := manage.NewDefaultManager()
	// token memory store
	//manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	//clientStore := store.NewClientStore()
	//clientStore.Set("000000", &models.Client{
	//	ID:     "000000",
	//	Secret: "999999",
	//	Domain: "http://localhost",
	//})
	oauthManager.MapClientStorage(&store.ClientStore{})

	initOAuthTokenStore()

	oauth.OAuth2Server = server.NewDefaultServer(oauthManager)
	oauth.OAuth2Server.SetAllowGetAccessRequest(false)
	oauth.OAuth2Server.SetClientInfoHandler(server.ClientFormHandler)

	oauth.OAuth2Server.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	oauth.OAuth2Server.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

}

func initOAuthTokenStore() {
	// use redis token store
	oauthManager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		DB:       15,
		Password: config.Conf.Redis.Password,
	}))
}
