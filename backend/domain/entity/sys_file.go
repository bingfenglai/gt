package entity

import "gorm.io/gorm"

type SysFile struct {
	gorm.Model
	OriginalName string `json:"original_name" gorm:"not null;size=64"`
	Uri          string `json:"uri" gorm:"not null;size=255"`
	Md5          string `json:"md5" gorm:"not null;size=32" `
	NameSpace    string `json:"name_space" gorm:"not null;size=64;default:gt"`
	State        string `json:"state" gorm:"not null;size=1;default:0"`
	TenantId     uint   `json:"tenant_id"`
}

func (receiver *SysFile) TableName() string {
	return "tb_sys_file"
}
