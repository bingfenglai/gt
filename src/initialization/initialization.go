package initialization

func InitAll() {

	initLogConfig()
	initDbConfig()
	initCacheConfig()
	initStorage()
	initService()
	initOAuth2Server()

}
