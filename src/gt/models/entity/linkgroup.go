package entity

import "time"

type LinkGroup struct {
	Id        int64
	GroupName string
	CreateAt  time.Time
	CreateBy  int64
	UpdateAt  time.Time
	UpdateBy  int64
	Status    int
}
