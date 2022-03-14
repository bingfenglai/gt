package service

import (
	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/common/model/session"
	"github.com/go-oauth2/oauth2/v4/manage"
	"go.uber.org/zap"
)

type IUserSessionService interface {
	// 跟据用户标识获取用户会话信息
	GetSession(uid string) (*session.UserSessionInfo, error)

	Create(uid string) error

	CreateWithTenant(uid string, tenantId string) error

	CreateWithAccessToken(uid, tenantId, accessToken string) error
}

type userSessionServiceImpl struct {
}

func (svc *userSessionServiceImpl) GetSession(uid string) (*session.UserSessionInfo, error) {

	if uid == "" {
		return nil, errors.ErrUserIDCannotBeEmpty
	}

	us := session.UserSessionInfo{}
	if err := CacheService.Get(session.USER_SESSION_PREFIX+uid, &us); err != nil {
		zap.L().Error(err.Error())
		return nil, errors.ErrUserNotFound
	}

	return &us, nil
}

func (svc *userSessionServiceImpl) Create(uid string) error {

	return svc.CreateWithTenant(uid, "-")

}

func (svc *userSessionServiceImpl) CreateWithTenant(uid string, tenantId string) error {
	if tenantId == "" {
		return errors.ErrTenantIdCannotBeEmpty
	}

	return svc.CreateWithAccessToken(uid, tenantId, "-")

}

func (svc *userSessionServiceImpl) CreateWithAccessToken(uid, tenantId, accessToken string) error {

	if err := svc.createRoleCheck(uid, tenantId, accessToken); err != nil {
		return err
	}

	roles, err := RoleService.GetSessionRolesByUid(uid)

	if err != nil {
		return err
	}

	roleIds := make([]int, len(roles))

	for i, role := range roles {
		roleIds[i] = role.RoleId
	}

	apis, err := ApiSesvice.GetSessionApisByRoleIds(roleIds)

	if err != nil {
		return err
	}

	us, _ := session.NewUserSession(uid, tenantId, accessToken, roles, apis)

	err = CacheService.Set(session.USER_SESSION_PREFIX+uid, us, manage.DefaultPasswordTokenCfg.AccessTokenExp)

	return err
}

func (svc *userSessionServiceImpl) createRoleCheck(uid, tenantId, accessToken string) error {
	if uid == "" {
		return errors.ErrUserIDCannotBeEmpty
	}

	if tenantId == "" {
		return errors.ErrTenantIdCannotBeEmpty
	}

	if accessToken == "" {
		return errors.ErrAccessTokenCannotBeEmpty
	}

	return nil
}
