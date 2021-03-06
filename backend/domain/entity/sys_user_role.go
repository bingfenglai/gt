package entity

import "time"

type UserRole struct {
	UserId    int64 `gorm:"primaryKey;not null"`
	RoleId    int64 `gorm:"primaryKey;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Status int `gorm:"default:0"`
	CreatedBy time.Time
	UpdatedBy time.Time
	TenantId int

}


func (u *UserRole) TableName()string{
	return "tb_sys_user_role"
}