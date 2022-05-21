package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name      string `gorm:"not null;size=24"`
	Code      string `gorm:"not null;size=24"`
	CreatedBy int64
	UpdatedBy int64
	Remark    string `gorm:"size:64;"`
	Status    int    `gorm:"default:0"`
	TenantId  int
}

func (r *Role) TableName() string {
	return "tb_sys_role"
}
