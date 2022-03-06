package errors

import "errors"

// 集中定义错误
var (
	UserNotFound = errors.New("用户不存在")
)
