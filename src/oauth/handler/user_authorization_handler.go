package handler

import (
	"github.com/bingfenglai/gt/oauth/utils"
	"net/http"
)

func UserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	username, err := utils.GetCurrentUsername(r)

	if err!=nil {
		
	}
	return username,nil
}
