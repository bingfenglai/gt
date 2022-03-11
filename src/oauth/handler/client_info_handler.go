package handler

import (
	"net/http"
	"strings"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/errors"
	"go.uber.org/zap"
)
const clientParamsName = "client"

// 获取当前认证请求当中的client信息，即user要登录的应用
// 此函数获取请求头当中的Authorization字段 解密后得到client的相关信息
// 
func ClientInfoHandler(r *http.Request) (clientID, clientSecret string, err error){
	
	s, err :=getAuthorizationStr(r)

	if err!=nil {
		return "","",err
	}
	zap.L().Info("client",zap.Any("client",s))
	s ,err = helper.AesDecryptCFB(s,[]byte(config.Conf.Encrypt.AesKey))

	if err!=nil {
		zap.L().Error(err.Error())
		return "","",errors.ErrClientUnauthorized
	}
	
	return getClientInfo(s)

}


func getAuthorizationStr(r *http.Request) (string,error){
	clientInfo :=  r.FormValue(clientParamsName)

	if clientInfo!=""{
		return clientInfo,nil
	}

	return "",errors.ErrClientUnauthorized
}

func getClientInfo(s string) ( clientID, clientSecret string,err error){
	zap.L().Info(s)
	err = errors.ErrClientUnauthorized
	sp := strings.Split(s,"&")

	zap.L().Info("client",zap.Any("",sp))

	if sp==nil||len(sp)!=2{
		return clientID,clientSecret,err
	}

	clientID = sp[0]
	clientSecret = sp[1]

	return clientID,clientSecret,nil

}