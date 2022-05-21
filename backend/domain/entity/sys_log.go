package entity

import (
	"gorm.io/gorm"
	"time"
)

type SysLog struct {
	gorm.Model

	// 操作者id
	Uid uint

	// 日志类型 1登录日志 2操作日志
	LogType int8 `gorm:"size 1;not null"`
	// 访问路径
	Uri string
	// 请求方法
	HttpMethod string
	// 入参
	Params string

	// 响应
	Response string

	// 请求耗时
	Cost int8
	// 状态
	Status int
	// 访问时间
	AccessTime time.Time
	// 用户代理
	UserAgent string

	// 用户代理客户端
	Client string

	// 操作系统
	OperationSystem string

	// 设备类型 1 Mobile 2 Tablet 3 Desktop
	DeviceType int8

	// Ip地址
	Ip uint
	// 错误信息
	Err string

	//扩展字段
	Extend string
}

func (sysLog *SysLog) TableName() string {
	return "tb_sys_log"
}
