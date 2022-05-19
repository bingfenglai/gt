package service

import (
	"github.com/bingfenglai/gt/service_interfaces"
	"log"

	"github.com/bingfenglai/gt/common/model/cache"
)

var CacheService cache.Cache

var CaptchaService service_interfaces.ICaptchaService

var UserService service_interfaces.IUserService
var RoleService service_interfaces.IRoleService

var ApiService service_interfaces.IApiService

var PasswordEncodeService service_interfaces.IPasswordEncoder

var OAuthClientService service_interfaces.IOAuthClientService

var ShortCodeService service_interfaces.IShortCodeService

var ShortCodeLogService service_interfaces.IShortCodeLogService

var EmailService service_interfaces.IEmailService

var TenantService service_interfaces.ITenantService

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
