package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/pojo/response"
	"github.com/wenlng/go-captcha/captcha"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type CaptchaServiceImpl struct {
}

func (c *CaptchaServiceImpl) GetImagesBehavioralCaptcha() (response.CaptchaResponse, error) {
	capt := captcha.GetCaptcha()

	dots, b64, th64, key, err := capt.Generate()

	zap.L().Info("验证值：")
	fmt.Println(dots)
	cacheKey := config.Conf.Captcha.Prefix + key
	CacheService.SetWithJson(cacheKey, dots, config.Conf.Captcha.ValidityPeriod)

	return response.CaptchaResponse{CaptchaId: key, ImageBase64: b64, ThumbBase64: th64}, err
}

func (c *CaptchaServiceImpl) Verify(src, captchaId string) (bool, error) {
	captchaId = config.Conf.Captcha.Prefix + captchaId
	ok, s := CacheService.GetWithJson(captchaId)
	if !ok {
		return false, errors.New("验证码已过期")
	}

	var dct map[int]captcha.CharDot

	err := json.Unmarshal([]byte(s), &dct)

	zap.L().Info("解析服务器缓存数据：", zap.String("value", s))

	if err != nil {
		return false, err
	}
	chkRet := false

	zap.L().Info("用户输入：" + src)

	if len(src) >= len(dct)*2 {
		chkRet = true
		zap.L().Info("开始检测点位置")
		split := strings.Split(src, ",")
		for i, dot := range dct {
			j := i * 2
			k := i*2 + 1
			//sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[j]), 64)
			//sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[k]), 64)
			sx, _ := strconv.ParseFloat(split[j], 64)
			sy, _ := strconv.ParseFloat(split[k], 64)
			// 检测点位置
			zap.L().Info("sx: ", zap.Float64("sx:", sx))
			zap.L().Info("sy: ", zap.Float64("sy", sy))
			zap.L().Info("dx: ", zap.Int("dx", dot.Dx))
			zap.L().Info("dy: ", zap.Int("dy", dot.Dy))
			chkRet = captcha.CheckPointDist(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height))
			if !chkRet {
				break
			}
		}
	}

	if chkRet {
		return true, nil
	} else {
		return false, errors.New("验证码错误")
	}

}
