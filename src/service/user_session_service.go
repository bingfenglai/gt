package service

import (
	"time"

	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/common/model/session"
	"go.uber.org/zap"
)

type IUserSessionService interface {
	// 跟据用户标识获取用户会话信息
	GetSession(uid string) (*session.UserSessionInfo, error)

	Create(uid string) error

	CreateWithTenant(uid string, tenantId string) error
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
	if uid == "" {
		return errors.ErrUserIDCannotBeEmpty
	}

	roles, err := RoleService.GetSessionRolesByUid(uid)

	if err != nil {
		return err
	}

	us, _ := session.NewUserSession(uid, tenantId, "", roles, nil)

	err = CacheService.Set(session.USER_SESSION_PREFIX+uid, us, time.Minute*30)

	return err

}
