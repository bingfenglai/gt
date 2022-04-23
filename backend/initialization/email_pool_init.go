package initialization

import (
	
	"net/smtp"
	

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/global"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
)

func initEmailPool() {

	zap.L().Info("配置邮件服务器")

	if global.EmailPool != nil {
		return
	}

	// if config.Conf.Email.Enable && config.Conf.Email.Auth == "" {
	// 	log.Default().Panicln("未配置邮箱服务器密钥，将从系统环境变量当中读取")
	// 	eauth := os.Getenv("EMAIL_AUTH")

	// 	if eauth != "" {
	// 		config.Conf.Email.Auth = eauth
	// 	} else {
	// 		panic("邮箱服务器密码未配置")
	// 	}
	// }

	auth := smtp.PlainAuth("", config.Conf.Email.SenderEmail,
		config.Conf.Email.Auth, config.Conf.Email.SmtpServerHost)
	emailPool, err := email.NewPool(config.Conf.Email.Address, 10, auth)
	if err != nil {
		//panic(err.Error())
		zap.L().Error("邮件配置出错", zap.Error(err))
	}
	emailPool.SetHelloHostname("127.0.0.1")

	global.EmailPool = emailPool
}
