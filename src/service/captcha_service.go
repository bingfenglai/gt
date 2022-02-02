package service

import "github.com/bingfenglai/gt/pojo/response"

type ICaptchaService interface {
	// 获取行为式图片验证码
	GetImagesBehavioralCaptcha() (response.CaptchaResponse, error)

	// 验证
	Verify(dots, captchaId string) (bool, error)
}
