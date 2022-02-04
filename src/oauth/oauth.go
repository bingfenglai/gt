package oauth

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"go.uber.org/zap"
)

var OAuth2Server *server.Server

func init() {
	if OAuth2Server == nil {
		zap.L().Error("oauth2 未初始化...")
	}
}
