package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bingfenglai/gt/common/constants"
	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/pojo/response"
	"github.com/google/uuid"
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
	CacheService.SetWithJson(cacheKey, dots, config.Conf.Captcha.ValidityPeriod*time.Minute)

	return response.CaptchaResponse{CaptchaId: key, ImageBase64: b64, ThumbBase64: th64}, err
}

func (c *CaptchaServiceImpl) ImagesBehavioralVerify(src, captchaId string) (bool, error) {
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

func (c *CaptchaServiceImpl) GetNumberCode(receiver string, channel int8) (captchaId string, err error) {

	if receiver == "" {
		return "", errors.New("接收验证码账号不能为空")
	}
	kp := config.Conf.Captcha.Prefix+receiver+"*"
	if flag,_ := CacheService.Keys(kp);flag{
		return "",errors.New("验证码发送太过频繁")
	}

	var code string

	switch channel {
	case constants.SEND_CODE_EMIAL:

		if err = helper.VerifyEmailFormat(receiver); err != nil {
			return
		}

		code, captchaId = c.genNumberCode(receiver)

		sec := config.Conf.Captcha.ValidityPeriod * time.Minute

		

		enText :="Verification code: " + code + " Valid within " + sec.String() + ".\n"

		zhText := "验证码："+code +" "+ sec.String() + "内有效"

		sendEmailParams := &params.EmailSimpleSendParams{
			Receivers: []string{receiver},
			Subject:   "[GT]Please verify your device",
			Text:      []byte(enText+zhText),
		}
		// 异步发送验证码
		if config.Conf.Server.Mode== constants.SERVER_MODE_TEST{
			EmailService.SendSimpleEmail(sendEmailParams)
		}else{
			go EmailService.SendSimpleEmail(sendEmailParams)
		}
		
		return

	}

	return "", errors.New("不支持该接收验证码渠道")
}

func (c *CaptchaServiceImpl) NumberCodeVerify(code string, captchaId string, receiver string) error {

	if config.Conf.Captcha.NumberCodeLength != len(code) {
		return errors.New("验证码不合法")
	}

	captchaId = config.Conf.Captcha.Prefix + receiver + ":" + captchaId

	if flag, jsonStr := CacheService.GetWithJson(captchaId); !flag || jsonStr == "" {

		return errors.New("验证码已过期")
	} else {
		c := ""
		json.Unmarshal([]byte(jsonStr), &c)

		if code != c {
			return errors.New("验证码已过期")
		}else{
			go CacheService.Delete(captchaId)
		}
	}

	return nil
}

func (c *CaptchaServiceImpl) genNumberCode(receiver string) (code, captchaId string) {
	l := config.Conf.Captcha.NumberCodeLength
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := l; i > 0; i-- {

		code = code + strconv.Itoa(r.Intn(9))
	}
	captchaId = uuid.NewString()

	CacheService.SetWithJson(config.Conf.Captcha.Prefix+receiver+":"+captchaId, code, config.Conf.Captcha.ValidityPeriod*time.Minute)
	return

}
