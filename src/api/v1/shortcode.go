package v1

import (
	"github.com/bingfenglai/gt/router"
	"net/http"

	"github.com/bingfenglai/gt/model/shortcodegen"
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/pojo/result"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

func GenShortCode(ctx *gin.Context){
	params :=params.GenShortCodeParams{}
	ctx.ShouldBindBodyWith(&params,binding.JSON)
	zap.L().Info("接收到参数",zap.Reflect("params",params))
	
	gen,_ := shortcodegen.GetShortCodeGeneratorByMethod(shortcodegen.Md5Gen)
	codes,_ :=gen.GenShortCode(params.OriginalLink)
	ctx.JSON(http.StatusOK,result.Ok(codes))

}


func init() {
	router.GetV1().POST("/shortCode/gen/:code",GenShortCode)

}