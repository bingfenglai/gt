package interfaces

import "github.com/bingfenglai/gt/domain/params"

// 邮件服务接口
type IEmailService interface {

	// 简要邮件发送方法，一般用于发送验证码等
	SendSimpleEmail(params *params.EmailSimpleSendParams) error
}
