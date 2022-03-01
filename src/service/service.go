package service

import (
	
	"log"

	"github.com/bingfenglai/gt/model/cache"
)

var CacheService cache.Cache

var CaptchaService ICaptchaService

var UserService IUserService

var PasswordEncodeService IPasswordEncoder

var AuthcService IAuthService


var ShortCodeService IShortCodeService

func InitService() {
	
	log.Default().Println("执行service初始化")
	CacheService = cache.GetCacheImpl()
	CaptchaService = &CaptchaServiceImpl{}
	UserService = &UserServiceImpl{}

	PasswordEncodeService = &PasswordEncoder{}

	AuthcService = &AuthenticationService{}


	ShortCodeService = &ShortCodeServiceImpl{}
}
