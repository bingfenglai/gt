package handler

import (
	"net/http"
	"strings"

	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/errors"
	"go.uber.org/zap"
)


const authorizationPrefix = "Basic "
const basicName = "Authorization"

// 获取当前请求当中的client信息，即user要登录的应用
// 此函数获取请求头当中的Authorization字段 解密后得到client的相关信息
// Authorization 前缀 Basic



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
	auth :=  r.Header.Get(basicName)
	if flag :=checkIsBasic(auth);flag{
		token := auth[len(authorizationPrefix):]
		
		return token,nil
	}

	return "",errors.ErrClientUnauthorized
}


func checkIsBasic(s string)bool{

	if s == "" {
		return false
	}

	return strings.HasPrefix(s,authorizationPrefix)
	
	
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