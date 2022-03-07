package errors

import "errors"

// 集中定义错误
var (
	ErrUserNotFound = errors.New("用户不存在")

	ErrAccountPasswordMismatch = errors.New("账号密码不匹配")
)
