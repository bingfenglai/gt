package utils

import (
	"github.com/bingfenglai/gt/common/model/session"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"time"
)

type GtContext struct {
	UserSession *session.UserSessionInfo
}

func NewGtContext(ctx *gin.Context) (context.Context, error) {
	user, err := GetCurrentUser(ctx.Request)
	if err != nil {
		return nil, err
	}
	return &GtContext{UserSession: user}, nil
}

func (*GtContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*GtContext) Done() <-chan struct{} {
	return nil
}

func (*GtContext) Err() error {
	return nil
}

func (*GtContext) Value(key interface{}) interface{} {
	return nil
}
