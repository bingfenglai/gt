package test

import (
	"testing"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/service"
	
	"go.uber.org/zap"
)

func TestSendEmail(t *testing.T) {

	p := &params.EmailSimpleSendParams{
		Receivers: []string{"1904454128@qq.com"},
		Subject:   "Testing",
		Text:      []byte("Testing"),
		HTML:      nil,
	}
	zap.L().Info("邮件配置",zap.Any("email_config",config.Conf.Email))
	if err := service.EmailService.SendSimpleEmail(p);err!=nil{
		t.Error(err.Error())
	}else {
		t.Log("邮件发送成功")
	}

}
