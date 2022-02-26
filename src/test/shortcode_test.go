package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bingfenglai/gt/model/shortcodegen"
	"github.com/bingfenglai/gt/service"
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


func TestCreate(t *testing.T){
	
	fmt.Println("testing")
	if sc,err :=service.ShortCodeService.CreateShortCode("https://www.google.com",false,true);err!=nil{
		t.Log(err.Error())
	}else{
		t.Log("shortcode :",sc.ShortCode)
		zap.L().Info("短码",zap.String("短码",sc.ShortCode))

	}





}