package test

import (
	"testing"

	"github.com/bingfenglai/gt/common/model/session"
	"go.uber.org/zap"
)

func TestGetUserSession(t *testing.T){

	if info,err :=session.UserSessionService.GetSession("3");err==nil{
		zap.L().Info("user session info",zap.Any("",info))
	}
}