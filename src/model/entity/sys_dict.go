package entity

import "github.com/jinzhu/gorm"

// 数据字典
type Dict struct {
	gorm.Model
	Name string `gorm:"not null;size:24"`
	Code string `gorm:"not null;size:24"`
	DictType int `gorm:"not null;default:0;comment:'字典类型0系统1业务'"`
	Status int `gorm:"default:0"`
	Remark string `gorm:"not null;size:64"`
	TenantId int

	
}

func (d *Dict) TableName()string{
	return "tb_sys_dict"
}