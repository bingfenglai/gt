package params

import "errors"

type TenantCreateParams struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
	ParentId int `json:"-"`
}

func (p *TenantCreateParams) Check() error {
	if p.Name != "" {
		return nil
	} else {
		return errors.New("租户名称不能为空")
	}
}