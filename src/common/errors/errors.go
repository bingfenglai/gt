package errors

import "errors"

// 集中定义错误
var (
	ErrUserNotFound = errors.New("用户不存在")

	ErrUserIDCannotBeEmpty = errors.New("用户id不能为空")

	ErrAccountPasswordMismatch = errors.New("账号密码不匹配")

	ErrAccessTokenCannotBeEmpty = errors.New("token参数不能为空")

	ErrClientUnauthorized = errors.New("客户端未授权")

	ErrParamsNotNull = errors.New("参数不能为空")

	ErrEmailContentIsNull = errors.New("邮件内容不能为空")

	ErrEmailFormat = errors.New("电子邮箱格式错误")

	ErrEmailCodeInvalid = errors.New("邮箱验证码不合法")

	ErrInvalidHttpMethod = errors.New("当前方法不被允许")
)
