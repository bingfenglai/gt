package test

import (
	"strings"
	"testing"

	"github.com/bingfenglai/gt/model/shortcodegen"
	"go.uber.org/zap"
)

var gen shortcodegen.ShortCodeGenerator

func TestGen(t *testing.T){
	var err error
	gen, err= shortcodegen.GetShortCodeGeneratorByMethod(shortcodegen.Md5Gen)

	if err!=nil {
		zap.L().Error(err.Error())
		t.Fail()
		return
	}

	codes,err :=gen.GenShortCode("https://www.google.com")

	if err!=nil {
		zap.L().Error(err.Error())
		t.Fail()
		return
	}

	zap.L().Info(strings.Join(codes,","))



	
}