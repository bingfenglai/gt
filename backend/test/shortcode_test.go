package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bingfenglai/gt/common/model/shortcodegen"
	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
)

var gen shortcodegen.ShortCodeGenerator

func TestGen(t *testing.T) {
	var err error
	gen, err = shortcodegen.GetShortCodeGeneratorByMethod(shortcodegen.Md5Gen)

	if err != nil {
		zap.L().Error(err.Error())
		t.Fail()
		return
	}

	codes, err := gen.GenShortCode("https://www.google.com")

	if err != nil {
		zap.L().Error(err.Error())
		t.Fail()
		return
	}

	zap.L().Info(strings.Join(codes, ","))

}

func TestCreate(t *testing.T) {

	fmt.Println("testing")
	if sc, err := service.ShortCodeService.CreateShortCode("https://www.baidu.com", false, true); err != nil {
		t.Log(err.Error())
	} else {

		zap.L().Info("短码", zap.String("短码", sc.ShortCode))

	}

}

func TestFindLinkByCode(t *testing.T) {
	if url, err := service.ShortCodeService.FindLinkByCode("JB29P5"); err != nil {
		zap.L().Error(err.Error())
	} else {
		zap.L().Info("url", zap.Any("", url))
	}

}
