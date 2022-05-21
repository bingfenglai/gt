package entity

import "gorm.io/gorm"

type Api struct {
	gorm.Model
	Name      string `gorm:"size:24;not null"`
	Uri       string `gorm:"size:36;not null"`
	Method    string `gorm:"size:36;not null"`
	Status    int    `gorm:"default:0"`
	Remark    string `gorm:"size:64;"`
	CreatedBy int64
	UpdatedBy int64
	TenantId  int
}

func (a *Api) TableName() string {
	return "tb_sys_api"
}
