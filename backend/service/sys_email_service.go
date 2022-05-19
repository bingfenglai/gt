package service

import (
	"time"

	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/global"
	"github.com/jordan-wright/email"
)

type emailServiceImpl struct {
}

func (e emailServiceImpl) SendSimpleEmail(params *params.EmailSimpleSendParams) error {
	if params == nil {
		return errors.ErrParamsNotNull
	}

	if err := params.Check(); err != nil {
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

	err := global.EmailPool.Send(email, time.Second*5)

	if err != nil && err.Error() == "timed out" {
		if err = global.EmailPool.Send(email, time.Second*10); err != nil && err.Error() == "timed out" {
			global.EmailPool.Send(email, time.Second*15)
		}
	}

	return err
}
