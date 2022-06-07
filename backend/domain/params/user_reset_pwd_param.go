package params

import "github.com/bingfenglai/gt/common/errors"

type ResetPwdParam struct {
	NewPwd string `json:"new_pwd"`
	Code   string `json:"code"`
}

func (receiver ResetPwdParam) Check() error {

	if receiver.NewPwd == "" {
		return errors.NewErrParamsNotNull("new_pwd")
	}

	if receiver.Code == "" {
		return errors.NewErrParamsNotNull("code")
	}

	return nil

}
