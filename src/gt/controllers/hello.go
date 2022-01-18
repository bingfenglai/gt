package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

type HelloController struct {
	web.Controller
}

func (c *HelloController) Get() {

	name := c.GetString("name")

	if name == "" {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.ResponseWriter.Write([]byte("缺少name参数"))
	}

	c.Ctx.ResponseWriter.Write([]byte("Hello " + name))
}
