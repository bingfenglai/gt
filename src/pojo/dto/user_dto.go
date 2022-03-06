package dto

import (
	"encoding/json"
	"github.com/bingfenglai/gt/model/entity"
)

type UserDTO struct {
	Username string
	Password string
	roles    []entity.Role
	apis     []entity.Api
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

func (i *UserDTO) UnmarshalBinary(data []byte) error  {
	return json.Unmarshal(data, &i)
}

