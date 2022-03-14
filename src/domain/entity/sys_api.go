package entity

import "github.com/jinzhu/gorm"

type Api struct {
	gorm.Model
	Name string `gorm:"size:24;not null"`
	Uri string 	`grom:"size:36;not null"`
	Method string `grom:"size:36;not null"`
	Status int `gorm:"default:0"`
	Remark string `grom:"size:64;"`
	CreatedBy int64
	UpdatedBy int64
	TenantId int
}


func(a *Api)TableName()string{
	return "tb_sys_api"
}