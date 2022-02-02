package response

type CaptchaResponse struct {
	CaptchaId string `json:"captcha_id"`
	// 主图
	ImageBase64 string `json:"image_url"`
	// 略缩图
	ThumbBase64 string `json:"thumb_base_64"`
}
