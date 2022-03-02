package entity

import (
	"time"

	"gorm.io/gorm"
)

// 短码链接转换访问记录
type ShortcodeLog struct {
	gorm.Model
	ShortcodeId uint64
	AccessTime time.Time
	UserAgent string
	OperationSystem string
	Client string
	Ip uint32
}


func (scl *ShortcodeLog) TableName()string{

	return "tb_biz_shortcode_log"
}
