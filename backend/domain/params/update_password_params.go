package params

import (
	"github.com/bingfenglai/gt/common/errors"
)

type UpdatePasswordParams struct {
	OldPwd string `json:"old_pwd" validate:"gt=0"`
	NewPwd string `json:"new_pwd" validate:"gt=0"`
}

func (receiver UpdatePasswordParams) Check() (err error){

	if receiver.OldPwd=="" {
		err = errors.ErrOldPwdIsNotNull
		return err
	}

	if receiver.NewPwd=="" {
		err = errors.ErrNewPwdIsNotNull
	}

	return
}
