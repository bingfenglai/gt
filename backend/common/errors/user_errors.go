package errors

var (
	ErrOldPwdIsNotNull = New("旧密码不能为空")

	ErrNewPwdIsNotNull = New("新密码不能为空")

	ErrUpdatedPwdLinkInvalid = New("重置密码链接已失效，请重新获取")
)
