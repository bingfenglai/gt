package service_interfaces

import (
	"mime/multipart"
)

type ISysFileService interface {
	Service
	UploadFile(file *multipart.FileHeader) error
}
