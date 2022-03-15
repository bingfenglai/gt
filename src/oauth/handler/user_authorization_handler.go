package handler

import (
	"net/http"

	"github.com/bingfenglai/gt/oauth/utils"
)

func UserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	username, err := utils.GetCurrentUId(r)

	// 内置代码中已有相关操作
	// if err!=nil {
	// 	w.Header().Add("Content-Type","application/json; charset=utf-8")
	// 	jsonResult, _ := json.Marshal(result.Fail(err.Error()))
	// 	http.Error(w, string(jsonResult),http.StatusFound)
	// }
	return username,err
}
