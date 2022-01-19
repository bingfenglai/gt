package entity

import "time"

type Link struct {
	Id           int64
	Md5          string
	ShortCode    string
	OriginalLink string
	CodeType     int
	CreateAt     time.Time
	CreateBy     int64
	UpdateAt     time.Time
	UpdateBy     int64
	Status       int
}
