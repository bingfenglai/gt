package service_interfaces

import "github.com/bingfenglai/gt/domain/params"

// IEmailService 邮件服务接口
type IEmailService interface {

	// SendSimpleEmail 简要邮件发送方法，一般用于发送验证码等
	SendSimpleEmail(params *params.EmailSimpleSendParams) error
}
