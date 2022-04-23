package entity

import "time"

type RoleApi struct {
	Role_id   int64 `gorm:"primaryKey"`
	Api_id    int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Status    int `gorm:"default:0"`
	CreatedBy time.Time
	UpdatedBy time.Time
	TenantId int
}

func (r *RoleApi) TableName() string {
	return "tb_sys_role_api"
}