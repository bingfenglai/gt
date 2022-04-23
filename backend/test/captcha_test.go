package test

import (
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/bingfenglai/gt/common/constants"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
)

func TestGetCaptcha(t *testing.T) {

	captcha, _ := service.CaptchaService.GetImagesBehavioralCaptcha()

	marshal, _ := json.Marshal(captcha)
	zap.L().Info("验证码:", zap.String("capt", string(marshal)))

}

func TestGetEmailCode(t *testing.T) {

	if c, err := service.CaptchaService.GetNumberCode("1904454128@qq.com", constants.SEND_CODE_EMIAL); err != nil {
		t.Error(err)

	} else {
		log.Default().Println("key: ", c)
	}

}

func TestVerityEmailCode(t *testing.T) {
	key := "824ac355-22ae-4d43-a760-81a14ff032b5"
	code := "836366"
	if err := service.CaptchaService.NumberCodeVerify(code, key, "1904454128@qq.com"); err != nil {
		t.Error(err)
	}
}

func TestGenNumberCode(t *testing.T) {
	l := config.Conf.Captcha.NumberCodeLength
	r := rand.New(rand.NewSource(time.Now().Unix()))
	code := ""
	for i := l; i > 0; i-- {

		code = code + strconv.Itoa(r.Intn(9))

	}
	log.Default().Println("生成的验证码: ", code)
}
