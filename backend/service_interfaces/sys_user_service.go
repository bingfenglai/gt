package service_interfaces

import (
	"context"
	"github.com/bingfenglai/gt/domain/dto"
	"github.com/bingfenglai/gt/domain/params"
)

type IUserService interface {
	FindUserByUsername(username string) (*dto.UserDTO, error)

	FindUserByUId(uid int) (*dto.UserDTO, error)

	// FindUserByUIdWithCache 先从缓存当中查找，找不到再走数据库并回写缓存
	FindUserByUIdWithCache(uid int) (*dto.UserDTO, error)

	FindUserByEmail(email string) (*dto.UserDTO, error)

	// FindUserByEmailWithRegister 跟据邮箱查找用户信息，如果不存在则进行注册
	FindUserByEmailWithRegister(email string) (*dto.UserDTO, error)

	// UpdatePwd 更新密码
	UpdatePwd(ctx context.Context, p *params.UpdatePasswordParams, uid int) error

	// UpdatePwdByCode 跟据code更新对应的密码，用于忘记密码的场景
	UpdatePwdByCode(ctx context.Context, param params.ResetPwdParam) error

	// SendUpdatePwdLink 发送密码重置邮件
	SendUpdatePwdLink(email string) error
}
