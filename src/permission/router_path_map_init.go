package permission

var handlerPathMap = make(map[string]string)

//func init() {
//	initHandlerPathMap()
//}

func GetRelativePath(handlerName string) string{

	if handlerPathMap ==nil||len(handlerPathMap)==0 {

	}

	return handlerPathMap[handlerName]
}

func SetHandlerNamePathMap(m map[string]string) {
	handlerPathMap = m
}