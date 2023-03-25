package service

import (
	"github.com/bingfenglai/gt/config"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
)

type SysFileServiceLocalImpl struct {
	baseService
}

func (s *SysFileServiceLocalImpl) UploadFile(file *multipart.FileHeader) (err error) {
	dns := config.Conf.FileConf.DefaultNameSpace
	os.Mkdir(dns, 666)
	uuid, _ := uuid.NewUUID()

	dst, err := os.Create(dns + "/" + uuid.String() + file.Filename)

	if err != nil {
		return
	}
	src, err := file.Open()

	if err != nil {
		return
	}

	_, err = io.Copy(dst, src)

	defer src.Close()

	return
}
