package params

import "errors"

type TenantCreateParams struct {
	Name   string
	Remark string
}

func (p *TenantCreateParams) Check() error {
	if p.Name != "" {
		return nil
	} else {
		return errors.New("租户名称不能为空")
	}
}