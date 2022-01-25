package entity

import "time"

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
	CreatedAt time.Time
	CreatedBy int64
	UpdatedAt time.Time
	UpdatedBy int64
	Status   int
}
