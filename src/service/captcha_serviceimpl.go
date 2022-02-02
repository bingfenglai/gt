package service

import (
	"errors"
	"fmt"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/pojo/response"
	"github.com/wenlng/go-captcha/captcha"
	"go.uber.org/zap"
)

type CaptchaServiceImpl struct {
}

func (c *CaptchaServiceImpl) GetImagesBehavioralCaptcha() (response.CaptchaResponse, error) {
	capt := captcha.GetCaptcha()

	dots, b64, th64, key, err := capt.Generate()

	zap.L().Info("验证值：")
	fmt.Println(dots)
	cacheKey := config.Conf.Captcha.Prefix + key
	CacheService.Set(cacheKey, dots, config.Conf.Captcha.ValidityPeriod)

	return response.CaptchaResponse{CaptchaId: key, ImageBase64: b64, ThumbBase64: th64}, err
}

func (c *CaptchaServiceImpl) Verify(dots, captchaId string) (bool, error) {
	captchaId = config.Conf.Captcha.Prefix + captchaId
	ok, s := CacheService.Get(captchaId)
	if !ok {
		return false, errors.New("验证码已过期")
	}

	if dots == s {
		return true, nil
	} else {
		return false, errors.New("验证码错误")
	}

}
