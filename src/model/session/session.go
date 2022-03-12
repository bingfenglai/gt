package session

import "github.com/bingfenglai/gt/pojo/dto"

// 用户会话信息
type UserSessionInfo struct {
	uid string
	tenantId string `json:"tenant_id"`
	roles []*dto.RoleSessionDTO
	apis []*dto.ApiSessionDTO
}

func NewUserSession()*UserSessionInfo