package params

// VerityCaptchaParams 验证码校验参数
type VerityCaptchaParams struct {
	Dots      string `json:"dots" binding:"required"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}
