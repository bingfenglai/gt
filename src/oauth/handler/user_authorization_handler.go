package handler

import (
	"github.com/bingfenglai/gt/oauth/utils"
	"net/http"
)

func UserAuthorizationHandler(_ http.ResponseWriter, r *http.Request) (userID string, err error) {
	username, err := utils.GetCurrentUsername(r)

	// 内置代码中已有相关操作
	//if err!=nil {
	//	jsonResult, _ := json.Marshal(result.Fail(err.Error()))
	//	http.Error(w, string(jsonResult),http.StatusFound)
	//}
	return username,err
}
