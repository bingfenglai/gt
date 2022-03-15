package handler

import (
	"context"
	"errors"
	"strconv"

	"github.com/bingfenglai/gt/service"

	"go.uber.org/zap"
)

// 用户名与密码方式认证授权时处理方式
func PasswordAuthorizationHandler(_ context.Context, username, password string) (userID string, err error) {

	zap.L().Info("当前登录用户", zap.String("username", username))
	user, err := service.UserService.FindUserByUId(username)

	if err != nil {
		return "", err
	}

	
	
	ok, err := service.PasswordEncodeService.Check(password, user.Password)

	if ok {
		return strconv.Itoa(user.Uid), nil
	}

	zap.L().Warn(err.Error())

	return "", errors.New("用户名或密码错误")

}
