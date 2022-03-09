package errors

import "errors"

// 集中定义错误
var (
	ErrUserNotFound = errors.New("用户不存在")

	ErrAccountPasswordMismatch = errors.New("账号密码不匹配")

	ErrClientUnauthorized = errors.New("客户端未授权")

	ErrParamsNotNull = errors.New("参数不能为空")

	ErrEmailContentIsNull = errors.New("邮件内容不能为空")
)
