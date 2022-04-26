package entity

import "gorm.io/gorm"

// 系统租户
type Tenant struct {
	gorm.Model
	ParentId uint `gorm:"default:0"`
	Name      string `gorm:"not null;size:24"`
	Code      string `gorm:"not null;size:24"`
	Status    int    `gorm:"default:0"`
	Remark    string `gorm:"size:64"`
	CreatedBy uint
	UpdatedBy uint
	Version uint `gorm:"default:0"`
}

func (t *Tenant) TableName() string {
	return "tb_sys_tenant"
}

func CreateTenant(tenantName, remark string) (tenant *Tenant) {
	tenant = &Tenant{
		Name:   tenantName,
		Remark: remark,
		Code:   "",
	}

	return
}
