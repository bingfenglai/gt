package interfaces

import "github.com/bingfenglai/gt/domain/response"

type ICaptchaService interface {
	// 获取行为式图片验证码
	GetImagesBehavioralCaptcha() (response.CaptchaResponse, error)

	// 验证
	ImagesBehavioralVerify(dots, captchaId string) (bool, error)

	// 获取数字验证码
	GetNumberCode(receiver string, channel int8) (captchaId string, err error)

	// 校验数字验证码
	NumberCodeVerify(code string, captchaId string, receiver string) error
}
