package v1

import (
	"net/http"

	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/router"
	"github.com/bingfenglai/gt/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Tags 租户API
// @Summary 创建租户
// @Description 创建租户
// @Success 200 {string} string  "ok"
// @Router /v1/tenant [post]
func CreateTenant(ctx *gin.Context) {

	p :=params.TenantCreateParams{}
	if err := ctx.ShouldBindBodyWith(&p,binding.JSON);err!=nil{
		ctx.JSON(http.StatusBadRequest,result.FailWithErr(err))
	}

	if err := p.Check();err!=nil {
		ctx.JSON(http.StatusBadRequest,result.FailWithErr(err))
	}

	if err := service.TenantService.Create(p,ctx);err!=nil {
		ctx.JSON(http.StatusBadRequest,err)
	}else{
		ctx.JSON(http.StatusOK,result.Ok(nil))
	}
}


// @Tags 租户API
// @Summary 获取租户列表
// @Description 获取租户列表
// @Success 200 {string} string  "ok"
// @Router /v1/tenant [post]
func TenantList(ctx *gin.Context){
	if list,err := service.TenantService.List();err==nil{
		ctx.JSON(http.StatusOK,result.Ok(list))
	}else{
		ctx.JSON(http.StatusBadRequest,result.Fail(err))
	}
}

func init()  {
	
	router.GetV1().POST("/tenant",CreateTenant)
	router.GetV1().GET("/tenant/list",TenantList)
}