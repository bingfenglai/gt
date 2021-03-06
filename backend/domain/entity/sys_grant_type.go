package entity

import "gorm.io/gorm"

// 系统授权模式
type OAuthGrantType struct {
	
	gorm.Model
	Name string `gorm:"not null;size:24"`
	Remark string `gorm:"not null;size:12"`
	Status int

}


func (ogt *OAuthGrantType) TableName()string{
	return "tb_sys_grant_type"
}