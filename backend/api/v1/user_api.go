package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/oauth/utils"
	"github.com/bingfenglai/gt/router"
	"github.com/bingfenglai/gt/service"
	"github.com/gin-gonic/gin"
)

// @Tags 用户API
// @Summary 获取个人信息
// @Description 获取个人信息
// @Success 200 {string} string  "ok"
// @Router /v1/shortCode [post]
func GetUserInfo(ctx *gin.Context) {

	if uid,err :=utils.GetCurrentUIdWithContext(ctx);err!=nil{
		ctx.JSON(http.StatusForbidden,result.Fail(err))
	}else{
		userDto,err := service.UserService.FindUserByUId(int(uid))
		if err!=nil {
			ctx.JSON(http.StatusBadRequest,result.Fail(err))
			return
		}
		ctx.JSON(http.StatusOK,result.Ok(userDto))
	}
}


func init(){

	router.GetV1().GET("/user/info",GetUserInfo)
}