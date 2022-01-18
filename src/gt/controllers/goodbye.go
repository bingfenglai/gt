package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
)

func Goodbye(ctx *context.Context) {

	ctx.ResponseWriter.Write([]byte("goodbye "))
}

func init() {
	beego.Get("/sayGoodbye", func(ctx *context.Context) {
		//name := ctx.Request.Form.Get("name")
		name := ctx.Input.Query("name")
		if name == "" {
			ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			ctx.ResponseWriter.Write([]byte("缺少name参数"))
			return
		}

		ctx.ResponseWriter.Write([]byte("goodbye " + name))
	})

}
