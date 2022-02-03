package service

import "github.com/bingfenglai/gt/model/cache"

var CacheService cache.Cache

var CaptchaService ICaptchaService

var UserService IUserService

var PasswordEncodeService IPasswordEncoder

var AuthcService IAuthService

func InitService() {
	CacheService = cache.GetCacheImpl()
	CaptchaService = &CaptchaServiceImpl{}
	UserService = &UserServiceImpl{}

	PasswordEncodeService = &PasswordEncoder{}

	AuthcService = &AuthenticationService{}
}
