package v1

import (
	"github.com/bingfenglai/gt/domain/params"
	"go.uber.org/zap"
	"net/http"
	"strconv"

	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/oauth/utils"
	"github.com/bingfenglai/gt/router"
	"github.com/bingfenglai/gt/service"
	"github.com/gin-gonic/gin"
)


// @Tags 用户API
// @Summary 获取个人信息
// @Description 获取个人信息
// @Accept  json
// @Produce  json
// @Success 200 {object} result.Result
// @Router /v1/user/info [post]
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


// @Tags 用户API
// @Summary 更改密码
// @Description 更改密码
// @Accept  json
// @Produce  json
// @Success 200 {object} result.Result
// @Router /v1/user/password [put]
func UpdatePassword(ctx *gin.Context) {

	zap.L().Debug("修改密码")
	p := params.UpdatePasswordParams{}
	err := ctx.BindJSON(&p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,result.Fail("参数绑定失败"))
		return
	}

	user, err := utils.GetCurrentUser(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized,result.Fail(err))
		return
	}
	uid, _ := strconv.Atoi(user.Uid)

	err = service.UserService.UpdatePwd(ctx, &p, uid)
	if err != nil {
		zap.L().Error("更新密码失败",zap.Error(err))
		ctx.JSON(http.StatusBadRequest,result.Fail(err.Error()))
	}else {
		ctx.JSON(http.StatusOK,result.Ok(nil))
	}

}


func init(){

	router.GetV1().GET("/user/info",GetUserInfo)
	router.GetV1().PUT("/user/password",UpdatePassword)
}