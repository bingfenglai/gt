package entity

import "github.com/jinzhu/gorm"

type Api struct {
	gorm.Model
	Name string `gorm:"size:24;not null"`
	HandlerName string `gorm:"size:128;not null"`
	HandlerMd5 string `gorm:"size:36;nut null"`
	Uri string 	`grom:"size:36;not null"`
	Methods string `grom:"size:36;not null"`
	Status int `gorm:"default:0"`
	Remark string `grom:"size:64;"`
	CreatedBy int64
	UpdatedBy int64
	TenantId int
}


func(a *Api)TableName()string{
	return "tb_sys_api"
}