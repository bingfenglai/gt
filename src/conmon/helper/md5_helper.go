package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMd5String(s string) string {

	md5 := md5.New().Sum([]byte(s))

	return hex.EncodeToString(md5)
}
