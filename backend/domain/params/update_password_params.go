package params

import "errors"

type UpdatePasswordParams struct {
	OldPwd string `json:"old_pwd"`
	NewPwd string `json:"new_pwd"`
}

func (receiver UpdatePasswordParams) Check() (err error){
	if receiver.OldPwd=="" {
		err = errors.New("");
	}
	


	return
}
