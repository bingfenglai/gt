package helper

import (
	"regexp"

	"github.com/bingfenglai/gt/common/errors"
)

// 校验邮箱格式

func VerifyEmailFormat(email string) error {
	//匹配电子邮箱
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(email) {
		return nil
	} 

	return errors.ErrEmailFormat
}
