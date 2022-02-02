package test

import (
	"encoding/json"
	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
	"testing"
)

func TestGetCaptcha(t *testing.T) {

	captcha, _ := service.CaptchaService.GetImagesBehavioralCaptcha()

	marshal, _ := json.Marshal(captcha)
	zap.L().Info("验证码:", zap.String("capt", string(marshal)))

}
