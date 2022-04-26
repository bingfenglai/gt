package dto

import (
	"encoding/json"
)

type UserDTO struct {
	Uid int `json:"uid"`
	TenantId int `json:"tenant_id"`
	Username string	`json:"username"`
	Password string `json:"-"`
}

//func (user *UserDTO) GetUsername() string {
//	return user.username
//}
//
//func (user *UserDTO) GetPassword() string {
//	return user.password
//}

// 实现缓存的编解码接口

func (i *UserDTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *UserDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &i)
}
