package service

import (
	"github.com/bingfenglai/gt/interfaces"
	"log"

	"github.com/bingfenglai/gt/common/model/cache"
)

var CacheService cache.Cache

var CaptchaService interfaces.ICaptchaService

var UserService interfaces.IUserService
var RoleService interfaces.IRoleService

var ApiService interfaces.IApiService

var PasswordEncodeService interfaces.IPasswordEncoder

var OAuthClientService interfaces.IOAuthClientService

var ShortCodeService interfaces.IShortCodeService

var ShortCodeLogService interfaces.IShortCodeLogService

var EmailService interfaces.IEmailService

var TenantService interfaces.ITenantService

func InitService() {

	log.Default().Println("执行service初始化")
	CacheService = cache.GetCacheImpl()
	CaptchaService = &captchaServiceImpl{}
	UserService = &userServiceImpl{}

	PasswordEncodeService = &passwordEncoder{}

	ShortCodeService = &shortCodeServiceImpl{}

	ShortCodeLogService = &shortCodeLogServiceImpl{}

	OAuthClientService = &oAuthClientServiceImpl{}

	EmailService = &emailServiceImpl{}

	RoleService = &roleServiceImpl{}

	ApiService = &apiServiceImpl{}

	TenantService = &tenantService{}
}
