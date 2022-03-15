package session

import (
	"encoding/json"

	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/domain/dto"
)

// 用户会话缓存前缀
const USER_SESSION_PREFIX = "session:"

// 用户会话信息
type UserSessionInfo struct {
	Uid         string `json:"uid"`
	TenantId    string `json:"tenant_id"`
	Roles       []*dto.RoleSessionDTO
	Apis        []*dto.ApiSessionDTO
	AccessToken string `json:"access_token"`
}

func NewUserSession(uid, tenantId, accessToken string, roles []*dto.RoleSessionDTO, apis []*dto.ApiSessionDTO) (*UserSessionInfo, error) {

	if uid == "" {
		return nil, errors.ErrUserIDCannotBeEmpty
	}

	if accessToken == "" {
		return nil, errors.ErrAccessTokenCannotBeEmpty
	}

	if roles == nil {
		roles = make([]*dto.RoleSessionDTO, 0)
	}

	if apis == nil {
		apis = make([]*dto.ApiSessionDTO, 0)
	}

	return &UserSessionInfo{
		Uid:         uid,
		TenantId:    tenantId,
		Roles:       roles,
		Apis:        apis,
		AccessToken: accessToken,
	}, nil
}

func (us *UserSessionInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(us)
}

func (us *UserSessionInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &us)
}
