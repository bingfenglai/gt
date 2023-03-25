package test

import (
	"io"
	"mime/multipart"
	"net/http"
	"testing"
)

const ossUploadFileUrl = "http://localhost:9527/v1/oss"

func TestOss(t *testing.T) {
	reader := getReader()
	http.Post(ossUploadFileUrl, "multipart/form-data", reader)

}

func getReader() io.Reader {

	reader := multipart.NewReader()

	return reader
}
