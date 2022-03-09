package test

import (
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/service"
	"testing"
)

func TestSendEmail(t *testing.T) {

	p := &params.EmailSimpleSendParams{
		Receivers: []string{"1904454128@qq.com"},
		Subject:   "Testing",
		Text:      []byte("Testing"),
		HTML:      nil,
	}
	if err := service.EmailService.SendSimpleEmail(p);err!=nil{
		t.Error(err.Error())
	}else {
		t.Log("邮件发送成功")
	}

}
