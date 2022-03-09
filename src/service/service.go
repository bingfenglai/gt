package service

import (
	
	"log"

	"github.com/bingfenglai/gt/model/cache"
)


type Service interface {
	Save(val interface{})(bool,error)
	DeleteById(id uint64)(bool,error)
}


var CacheService cache.Cache

var CaptchaService ICaptchaService

var UserService IUserService

var PasswordEncodeService IPasswordEncoder

var OAuthClientService IOAuthClientService


var ShortCodeService IShortCodeService

var ShortCodeLogService IShortCodeLogService

var EmailService IEmailService

func InitService() {
	
	log.Default().Println("执行service初始化")
	CacheService = cache.GetCacheImpl()
	CaptchaService = &CaptchaServiceImpl{}
	UserService = &UserServiceImpl{}

	PasswordEncodeService = &PasswordEncoder{}

	ShortCodeService = &ShortCodeServiceImpl{}

	ShortCodeLogService = &ShortCodeLogServiceImpl{}

	OAuthClientService = &OAuthClientServiceImpl{}

	EmailService = &EmailServiceImpl{}
}
