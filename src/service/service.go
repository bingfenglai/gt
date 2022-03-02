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

var AuthcService IAuthService


var ShortCodeService IShortCodeService

var ShortCodeLogService IShortCodeLogService

func InitService() {
	
	log.Default().Println("执行service初始化")
	CacheService = cache.GetCacheImpl()
	CaptchaService = &CaptchaServiceImpl{}
	UserService = &UserServiceImpl{}

	PasswordEncodeService = &PasswordEncoder{}

	AuthcService = &AuthenticationService{}


	ShortCodeService = &ShortCodeServiceImpl{}

	ShortCodeLogService = &ShortCodeLogServiceImpl{}
}
