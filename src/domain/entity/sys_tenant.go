package entity

import "gorm.io/gorm"

// 系统租户
type Tenant struct {
	gorm.Model
	Name string `gorm:"not null;size:24"`
	Code string `gorm:"not null;size:24"`
	Status int `gorm:"default:0"`
	Remark string `gorm:"not null;size:64"`
	CreatedBy int64
	UpdatedBy int64
	TenantId int
	
}


func(t *Tenant)TableName()string{
	return "tb_sys_tenant"
}