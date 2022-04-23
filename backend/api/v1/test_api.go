package v1

import (
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
	"net/http"

	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/router"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {

	name := ctx.Params.ByName("name")
	// age := ctx.Params.ByName("age")

	age := ctx.Request.FormValue("age")

	ctx.JSON(http.StatusOK, result.Ok(name+age))

}

func TestCtx(ctx *gin.Context) {
	user := entity.User{
		Username:  "zhangsan",
		Password:  "{noon}123",
		Email:     "qq@123.com",
		CreatedBy: 0,
		UpdatedBy: 0,
		Status:    0,
	}
	global.DB.WithContext(ctx).Create(&user)

	ctx.JSON(http.StatusOK, result.Ok(nil))

}

func init() {
	router.GetV1().GET("/test/:name/:age/print", Test)
	router.GetV1().GET("/test/ctx", TestCtx)
}
