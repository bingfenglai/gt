package v1

import (
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/bingfenglai/gt/router"
	"github.com/bingfenglai/gt/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"net/http"
)

func Login(ctx *gin.Context) {

}

func Logout(ctx *gin.Context) {

}

func Token(ctx *gin.Context) {
	p := params.AuthcParams{}

	if err := ctx.ShouldBindBodyWith(&p, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(err.Error()))
		return
	}

	authentication, err := service.AuthcService.Authentication(p)

	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(http.StatusOK, result.Fail(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Ok(authentication))

}

func RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.Query("refreshToken")

	if refreshToken == "" {
		ctx.JSON(http.StatusOK, result.Fail("刷新令牌不能为空"))
		return
	}

	p := params.AuthcParams{Principal: refreshToken, GrantType: "refreshToken"}
	authentication, err := service.AuthcService.Authentication(p)

	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Ok(authentication))

}

func init() {
	router.GetV1().POST("/oauth2/token", Token)
	router.GetV1().Any("/oauth2/refreshToken", RefreshToken)

}
