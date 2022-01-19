package models

import "time"

type Link struct {
	Id           int
	Md5          string
	ShortCode    string
	OriginalLink string
	CreateAt     time.Time
	UpdateAt     time.Time
	Status       int
}
