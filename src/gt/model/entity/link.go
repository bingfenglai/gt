package entity

import "time"

type Link struct {
	Id           int64
	Md5          string
	ShortCode    string
	OriginalLink string
	// 游客创建的code为临时code 0,注册用户创建的code为永久code 1
	CodeType int
	CreateAt time.Time
	CreateBy int64
	UpdateAt time.Time
	UpdateBy int64
	// 0启用 1未启用
	Status int
}
