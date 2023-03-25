package v1

import (
	"github.com/bingfenglai/gt/domain/result"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/router"
	"github.com/bingfenglai/gt/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"
)

const UploadFileKey = "file"

// @Tags 文件API
// @Summary 上传文件
// @Description 上传文件
// @Accept multipart/form-data
// @Success 200 {object} result.Result
// @Router /v1/oss [post]
func UploadFile(ctx *gin.Context) {

	file, fh, err := ctx.Request.FormFile(UploadFileKey)
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	if err != nil {
		global.Log.Error("接收文件错误", zap.Field{Key: "err", String: err.Error()})
		ctx.JSON(http.StatusOK, result.FailWithErr(err))
		return
	}
	filename := fh.Filename
	global.Log.Info("文件名称：", zap.String("file_name", filename))
	uuid, _ := uuid.NewUUID()
	filename = "temp/" + uuid.String() + "." + strings.Split(filename, ".")[1]
	os.Mkdir("temp", 666)
	err = ctx.SaveUploadedFile(fh, filename)

	if err != nil {
		global.Log.Error("保存文件错误", zap.Field{Key: "err", String: err.Error()})
		ctx.JSON(http.StatusOK, result.FailWithErr(err))
		return
	}

	ctx.JSON(http.StatusOK, result.Ok(nil))

}

// @Tags 文件API
// @Summary 上传文件到本地
// @Description 上传文件到本地
// @Param file 文件
// @Accept multipart/form-data
// @Success 200 {object} result.Result
// @Router /v1/oss [post]
func UploadFileWithLocal(ctx *gin.Context) {
	_, fh, err := ctx.Request.FormFile(UploadFileKey)

	if err != nil {
		global.Log.Error("接收文件错误", zap.Field{Key: "err", String: err.Error()})
		ctx.JSON(http.StatusOK, result.FailWithErr(err))
		return
	}

	err = service.SysFileServie.UploadFile(fh)
	if err != nil {
		global.Log.Error("接收文件错误", zap.Field{Key: "err", String: err.Error()})
		ctx.JSON(http.StatusOK, result.FailWithErr(err))
		return
	}

	ctx.JSON(http.StatusOK, result.Ok(nil))

}

func init() {
	router.GetV1().POST("/oss", UploadFile)
	router.GetV1().POST("/oss/local", UploadFileWithLocal)
}
