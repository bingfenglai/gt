package utils

import (
	"context"
	"errors"
	"github.com/bingfenglai/gt/common/model/session"
	gtContext "github.com/bingfenglai/gt/context"
	"github.com/bingfenglai/gt/oauth"
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/bingfenglai/gt/service"

	"net/http"
)

// 获取当前会话当中的user信息
func GetCurrentUser(req *http.Request) (*session.UserSessionInfo, error) {
	uid, err := GetCurrentUId(req)
	if err != nil {
		return nil, err
	}

	return service.UserSessionService.GetSession(uid)

}

func GetCurrentUId(req *http.Request) (string, error) {

	tokenInfo, err := oauth.OAuth2Server.ValidationBearerToken(req)

	if err != nil {
		return "", err
	}

	userID := tokenInfo.GetUserID()

	return userID, nil
}

func GetCurrentUIdWithContext(ctx context.Context) (uid int64, err error) {

	switch ctx.(type) {
	case *gtContext.GtContext:
		gtCtx := ctx.(*gtContext.GtContext)
		uid = gtCtx.Value("uid").(int64)
		return
	case *gin.Context:
		ginCtx := ctx.(*gin.Context)
		id, err := GetCurrentUId(ginCtx.Request)
		if err != nil {
			return 0, err
		}
		idInt, err := strconv.Atoi(id)
		uid = int64(idInt)
		return uid,nil
	}

	return uid, errors.New("context error")
}

func GetCurrentUserWithContext(ctx context.Context) (user *session.UserSessionInfo, err error) {
	uid, err := GetCurrentUIdWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return service.UserSessionService.GetSession(strconv.FormatInt(uid, 10))
}
