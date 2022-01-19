package entity

import "time"

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
	CreateAt time.Time
	CreateBy int64
	UpdateAt time.Time
	UpdateBy int64
	Status   int
}
