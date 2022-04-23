package oauth

import (
	"github.com/bingfenglai/gt/oauth/server"
	"go.uber.org/zap"
)

var OAuth2Server *server.CustomOAuthServer


func init() {
	if OAuth2Server == nil {
		zap.L().Error("oauth2 未初始化...")
	}

}
