package initialization

func InitAll() {

	initLogConfig()
	initDbConfig()
	initCacheConfig()
	initService()
	initOAuth2Server()

}
