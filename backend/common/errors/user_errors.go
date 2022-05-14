package errors

var (
	ErrOldPwdIsNotNull = New("旧密码不能为空")

	ErrNewPwdIsNotNull = New("新密码不能为空")
)

