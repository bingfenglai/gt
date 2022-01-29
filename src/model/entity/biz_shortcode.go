package entity

import "github.com/jinzhu/gorm"

type ShortCode struct {
	gorm.Model
	GroupId int64
	Md5          string
	ShortCode    string
	Original string
	// 游客创建的code为临时code 0,注册用户创建的code为永久code 1
	CodeType int

	CreatedBy int64

	UpdatedBy int64
	// 0启用 1未启用
	Status int `gorm:"default:0"`

	TenantId int
}

func (l *ShortCode) TableName() string {
	return "tb_biz_link"
}