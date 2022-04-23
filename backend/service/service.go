package service

import (
	"log"
	
	"github.com/bingfenglai/gt/common/model/cache"
)

type Service interface {
	Save(val interface{}) (bool, error)
	DeleteById(id uint64) (bool, error)
}

var CacheService cache.Cache

var CaptchaService ICaptchaService

var UserService IUserService

var UserSessionService IUserSessionService

var RoleService IRoleService

var ApiSesvice IApiService

var PasswordEncodeService IPasswordEncoder

var OAuthClientService IOAuthClientService

var ShortCodeService IShortCodeService

var ShortCodeLogService IShortCodeLogService

var EmailService IEmailService

var TenantService ITenantService

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

	UserSessionService = &userSessionServiceImpl{}

	RoleService = &roleServiceImpl{}

	ApiSesvice = &apiServiceImpl{}

	TenantService = &tenantService{}
}
