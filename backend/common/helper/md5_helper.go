package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMd5String(s string) string {

	md5 := md5.New().Sum([]byte(s))

	return hex.EncodeToString(md5)
}

func ToMd5String32(s string) string {

	hash := md5.New()

	hash.Write([]byte(s))

	return hex.EncodeToString(hash.Sum(nil))
}
