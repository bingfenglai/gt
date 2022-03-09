package handler

import (
	
	"github.com/bingfenglai/gt/service"
	"github.com/go-oauth2/oauth2/v4"
)

func ClientAuthorizedHandler(clientID string, grant oauth2.GrantType) (allowed bool, err error) {

	
	return service.OAuthClientService.CheckGrantType(clientID,string(grant))
	
}