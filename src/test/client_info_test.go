package test

import (
	
	"testing"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/config"
)

func TestAesEncode(t *testing.T){
	s := "999&999"
	if enc, err :=helper.AesEncryptCFB([]byte(s),[]byte(config.Conf.Encrypt.AesKey));err!=nil{
	
		t.Error(err)
	}else {
		
		t.Log("密文：",enc)
		
		if news, err := helper.AesDecryptCFB(enc,[]byte(config.Conf.Encrypt.AesKey));err!=nil{
			t.Error(err)
			

		}else{
			t.Log("解密后：",news)
		}
	}

}