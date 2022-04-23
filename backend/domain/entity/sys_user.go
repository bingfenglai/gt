package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;size=24"`
	Password string `gorm:"not null;size=128"`
	Email    string `gorm:"not null;size=64"`

	CreatedBy int64
	UpdatedBy int64
	Status    int `gorm:"default:0"`
	TenantId  int
}

func (user *User) TableName() string {
	return "tb_sys_user"
}
