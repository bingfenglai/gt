package utils

import (

	"github.com/bingfenglai/gt/oauth"
	"github.com/bingfenglai/gt/pojo/dto"
	"github.com/bingfenglai/gt/service"
	"github.com/go-oauth2/oauth2/v4/errors"

	"net/http"
)

// 获取当前会话当中的user信息
func GetCurrentUser(req *http.Request) (*dto.UserDTO,error)  {
	username, err := GetCurrentUsername(req)
	if err != nil {
		return nil,err
	}

	return service.UserService.FindUserByUsernameWithCache(username)

}

func GetCurrentUsername(req *http.Request) (string,error) {
	var token string
	var flag bool
	if token,flag =oauth.OAuth2Server.BearerAuth(req);!flag{
		return "",errors.ErrInvalidAccessToken
	}

	tokenInfo, err := oauth.OAuth2Server.Manager.LoadAccessToken(req.Context(), token)

	if err != nil {
		return "",err
	}

	userID := tokenInfo.GetUserID()

	return userID,nil
}
