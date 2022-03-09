package service

import (
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/errors"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/jordan-wright/email"
	"time"
)

// 邮件服务接口
type IEmailService interface {

	// 简要邮件发送方法，一般用于发送验证码等
	SendSimpleEmail(params *params.EmailSimpleSendParams)error
}

type EmailServiceImpl struct {

}

func (e EmailServiceImpl) SendSimpleEmail(params *params.EmailSimpleSendParams)error{
	if params==nil {
		return errors.ErrParamsNotNull
	}

	if err := params.Check();err!=nil{
		return err
	}
	
	email := &email.Email{
		ReplyTo:     nil,
		From:        config.Conf.Email.SenderEmail,
		To:          params.Receivers,
		Bcc:         nil,
		Cc:          nil,
		Subject:     params.Subject,
		Text:        params.Text,
		HTML:        params.HTML,
		Sender:      config.Conf.Email.SenderEmail,
		Headers:     nil,
		Attachments: nil,
		ReadReceipt: nil,
	}

	return global.EmailPool.Send(email, time.Second)
}
