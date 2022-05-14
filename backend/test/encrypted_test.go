package test

import (
	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/config"
	"log"
	"testing"
)


func TestEncrypted(t *testing.T) {

	s1, err := helper.AesEncryptCFB([]byte("123456"), []byte(config.Conf.Encrypt.AesKey));
	s2, err := helper.AesEncryptCFB([]byte("123456789"), []byte(config.Conf.Encrypt.AesKey));
	if err != nil {
		t.Error(err)
	}else {
		log.Default().Println(s1)
		log.Default().Println(s2)
	}

}
