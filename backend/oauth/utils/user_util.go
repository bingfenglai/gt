package utils

import (
	"context"
	"errors"
	"strconv"

	"github.com/bingfenglai/gt/common/model/session"
	"github.com/bingfenglai/gt/oauth"
	"github.com/gin-gonic/gin"

	"net/http"
)

// 获取当前会话当中的user信息
func GetCurrentUser(req *http.Request) (*session.UserSessionInfo, error) {
	uid, err := GetCurrentUId(req)
	if err != nil {
		return nil, err
	}

	return session.UserSessionService.GetSession(uid)

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
	case *GtContext:
		gtCtx := ctx.(*GtContext)
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
		return uid, nil
	}

	return uid, errors.New("context error")
}

func GetCurrentUserWithContext(ctx context.Context) (user *session.UserSessionInfo, err error) {
	uid, err := GetCurrentUIdWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return session.UserSessionService.GetSession(strconv.FormatInt(uid, 10))
}

func GetCurrentTenantId(ctx context.Context)(tenantId int,err error){
	if user,err := GetCurrentUserWithContext(ctx);err==nil&&user.TenantId!="-"{
		
		return strconv.Atoi(user.TenantId)
		
	}
	return -1,err
}
