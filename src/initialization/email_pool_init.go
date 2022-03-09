package initialization

import (
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/global"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"net/smtp"
)

func initEmailPool() {

	zap.L().Info("配置邮件服务器")

	if global.EmailPool !=nil {
		return
	}

	auth := smtp.PlainAuth("",config.Conf.Email.SenderEmail,
		config.Conf.Email.Auth,config.Conf.Email.SmtpServerHost)
	emailPool, err := email.NewPool(config.Conf.Email.Address, 10, auth)
	if err!=nil {
		//panic(err.Error())
		zap.L().Error("邮件配置出错",zap.Error(err))
	}
	emailPool.SetHelloHostname("127.0.0.1")

	global.EmailPool = emailPool
}
