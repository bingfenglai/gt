package test

import (
	"log"
	"testing"

	"github.com/bingfenglai/gt/service"
)

func TestEncode(t *testing.T){
	s,_ :=service.PasswordEncodeService.Encode("123456")

	log.Default().Println("加密后的密码： ",s)
	log.Default().Println("加密后的密码长度： ",len(s))

	flag,err :=service.PasswordEncodeService.Check("123456",s)

	log.Default().Println(flag,err)
	
}