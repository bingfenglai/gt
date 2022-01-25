package entity

import "time"



// GORM 倾向于约定，而不是配置。
//默认情况下，GORM 使用 ID 作为主键，使用结构体名的 蛇形复数 作为表名，字段名的 蛇形 作为列名
//，并使用 CreatedAt、UpdatedAt 字段追踪创建、更新时间
type LinkGroup struct {
	Id        int64  `gorm:"autoIncrement"`
	GroupName string `gorm:"size:12;not null"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`
	CreatedBy  int64 `gorm:"not null"`
	UpdatedAt  time.Time
	UpdatedBy  int64
	Status    int `gorm:"not null;default:0"`
}

func (l *LinkGroup)TableName() string{
	return "tb_link_group"
}
