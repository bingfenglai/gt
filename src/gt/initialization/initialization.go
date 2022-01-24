package initialization

func InitAll() {

	initLogConfig()
	initDbConfig()
	initCacheConfig()
	initService()

}
