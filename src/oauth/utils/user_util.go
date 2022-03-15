package utils

import (
	"github.com/bingfenglai/gt/common/model/session"
	"github.com/bingfenglai/gt/oauth"

	"github.com/bingfenglai/gt/service"

	"net/http"
)

// 获取当前会话当中的user信息
func GetCurrentUser(req *http.Request) (*session.UserSessionInfo, error) {
	uid, err := GetCurrentUId(req)
	if err != nil {
		return nil, err
	}

	return service.UserSessionService.GetSession(uid)

}

func GetCurrentUId(req *http.Request) (string, error) {

	tokenInfo, err := oauth.OAuth2Server.ValidationBearerToken(req)

	if err != nil {
		return "", err
	}

	userID := tokenInfo.GetUserID()

	return userID, nil
}
