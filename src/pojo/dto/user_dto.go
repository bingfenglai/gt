package dto

import "github.com/bingfenglai/gt/model/entity"

type UserDTO struct {
	Username string
	Password string
	roles    []entity.Role
	apis     []entity.Api
}
