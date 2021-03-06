package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bingfenglai/gt/domain/result"
)

// 自定义token 响应格式
func ResponseTokenHandler(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error {

	var r result.Result

	errFlag := true

	if data != nil && data["error"] != nil {
		msg := data["error_description"]
		m := msg.(string)
		r = result.FailWithMsg(m, nil)
		data["error_code"] = 0
		errFlag = false
		w.WriteHeader(400)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if errFlag {
		r = result.Ok(data)
	}
	jsonByte, _ := json.Marshal(r)
	w.Write(jsonByte)

	return nil
}
