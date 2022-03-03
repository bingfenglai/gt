package entity

import "time"

// 客户端授权模式
type ClientGrantType struct {

	ClientId uint64
	ClientBizId string `gorm:"not null;size:32"`
	GrantTypeId uint64
	Status int
	CreatedBy uint64
	CreatedAt time.Time
	UpdatedBy uint64
	UpdatedAt time.Time
	DeletedAt time.Time  
	
}


func (cgt *ClientGrantType) TableName()string{

	return "tb_sys_client_grant_type"
}