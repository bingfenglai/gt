package entity

import "gorm.io/gorm"

// 字典项
type DictItem struct {
	gorm.Model
	DictId int `gorm:"not null"`
	Name string `gorm:"not null;size:24"`
	Code string `gorm:"not null;size:24"`
	Value string `gorm:"not null;size:24"`
	Status int `gorm:"default:0"`
	Remark string `gorm:"not null;size:64"`
	CreatedBy int64
	UpdatedBy int64
	TenantId int
}


func (d *DictItem)TableName()string{
	return "tb_sys_dict_item"
}