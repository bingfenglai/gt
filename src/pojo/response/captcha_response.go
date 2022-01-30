package response

type CaptchaResponse struct {
	CaptchaId string `json:"captcha_id"`
	ImageUrl string `json:"image_url"`
}
