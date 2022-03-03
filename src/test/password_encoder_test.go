package test

import (
	"log"
	"testing"

	"github.com/bingfenglai/gt/service"
)

func TestEncode(t *testing.T){
	s,_ :=service.PasswordEncodeService.Encode("123456")

	log.Default().Println("加密后的密码： ",s)
	
}