package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/pojo/result"
	"github.com/bingfenglai/gt/router"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context){

	

	name := ctx.Params.ByName("name")
	// age := ctx.Params.ByName("age")

	age := ctx.Request.FormValue("age")

	ctx.JSON(http.StatusOK,result.Ok(name+age))


}


func init(){
	router.GetV1().GET("/test/:name/print",Test)
}