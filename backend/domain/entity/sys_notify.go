package entity

import (
	"gorm.io/gorm"
)

type SysNotify struct {
	gorm.Model
	// 通知渠道
	channel int8
}
