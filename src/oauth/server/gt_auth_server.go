package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bingfenglai/gt/oauth/manager"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/server"
	"go.uber.org/zap"
	
)

// 增强server
type CustomOAuthServer struct {
	Server                       *server.Server
	CustomOAuthManager           manager.CustomOAuthManager
	emailVerificationCodeHandler EmailVerificationCodeHandler
}

// 工厂方法
func NewDefaultCustomOAuthServer(orignManager oauth2.Manager) *CustomOAuthServer {

	return &CustomOAuthServer{
		Server: server.NewDefaultServer(orignManager),
		CustomOAuthManager: manager.NewCustomOAuthManager(orignManager),
	}
}

func (s *CustomOAuthServer) SetEmailVerificationCodeHandler(handler EmailVerificationCodeHandler) {
	s.emailVerificationCodeHandler = handler
}

// HandleTokenRequest token request handling
func (s *CustomOAuthServer) HandleTokenRequest(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	gt, tgr, err := s.ValidationTokenRequest(r)
	if err != nil {
		return s.tokenError(w, err)
	}

	ti, err := s.GetAccessToken(ctx, gt, tgr)
	if err != nil {
		return s.tokenError(w, err)
	}

	return s.token(w, s.Server.GetTokenData(ti), nil)
}

// 认证流程
func (s *CustomOAuthServer) ValidationTokenRequest(r *http.Request) (oauth2.GrantType, *oauth2.TokenGenerateRequest, error) {
	zap.L().Info("走代理认证流程")

	if v := r.Method; !(v == "POST" ||
		(s.Server.Config.AllowGetAccessRequest && v == "GET")) {
		return "", nil, errors.ErrInvalidRequest
	}
	gtStr := r.FormValue("grant_type")
	gt := oauth2.GrantType(gtStr)
	// gt := gtStr.(oauth2.GrantType)
	zap.L().Info("当前认证模式为", zap.Any("gtStr: ", gtStr))

	
	// gt := getGrantTypeByName(gtStr)

	zap.L().Info("当前认证模式为", zap.Any("gt: ", gt))

	if !s.CheckGrantType(gt) {
		return s.Server.ValidationTokenRequest(r)
	}

	clientID, clientSecret, err := s.Server.ClientInfoHandler(r)
	if err != nil {
		return "", nil, err
	}

	tgr := &oauth2.TokenGenerateRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Request:      r,
	}
	zap.L().Info("当前认证模式为", zap.Any("gt: ", gt))

	switch gt {
	case EmailCode:

		if s.emailVerificationCodeHandler == nil {
			return "", nil, errors.ErrUnsupportedGrantType
		}

		userId, err := s.emailVerificationCodeHandler(r)
		if err != nil {
			return "", nil, err
		}

		tgr.UserID = userId

		// 使其等同于密码登录的令牌配置，简化开发
		// gt = oauth2.PasswordCredentials
	default:
		// 原始oauth2 支持的认证模式
		return s.Server.ValidationTokenRequest(r)
	}

	return gt, tgr, nil

}

// 处理 获取访问令牌错误
func (s *CustomOAuthServer) tokenError(w http.ResponseWriter, err error) error {
	data, statusCode, header := s.Server.GetErrorData(err)
	return s.token(w, data, header, statusCode)
}

// 响应token数据
func (s *CustomOAuthServer) token(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error {
	if fn := s.Server.ResponseTokenHandler; fn != nil {
		return fn(w, data, header, statusCode...)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")

	for key := range header {
		w.Header().Set(key, header.Get(key))
	}

	status := http.StatusOK
	if len(statusCode) > 0 && statusCode[0] > 0 {
		status = statusCode[0]
	}

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// GetAccessToken access token
// 跟据gt 判断认证方式是否为自定义方式，不是则走默认的
func (s *CustomOAuthServer) GetAccessToken(ctx context.Context, gt oauth2.GrantType, tgr *oauth2.TokenGenerateRequest) (oauth2.TokenInfo, error) {
	
	// 判断是否为扩展的认证模式
	if !s.CheckGrantType(gt){
		return s.Server.GetAccessToken(ctx,gt,tgr)
	}

	if fn := s.Server.ClientAuthorizedHandler; fn != nil {
		allowed, err := fn(tgr.ClientID, gt)
		if err != nil {
			return nil, err
		} else if !allowed {
			return nil, errors.ErrUnauthorizedClient
		}
	}

	switch gt {
	case EmailCode:
		if fn := s.Server.ClientScopeHandler; fn != nil {
			allowed, err := fn(tgr)
			if err != nil {
				return nil, err
			} else if !allowed {
				return nil, errors.ErrInvalidScope
			}
		}


		return s.CustomOAuthManager.GenerateAccessToken(ctx,gt,tgr)

	}

	return nil,errors.ErrUnsupportedGrantType

}

func (s *CustomOAuthServer) HandleAuthorizeRequest(w http.ResponseWriter, r *http.Request) error {
	return s.Server.HandleAuthorizeRequest(w, r)
}

func (s *CustomOAuthServer) CheckGrantType(gt oauth2.GrantType) bool {

	return CheckGrantType(gt)

}


func CheckGrantType(gt oauth2.GrantType) bool {

	for _, v := range customGrantTypes {
		if v == gt {
			return true
		}
	}

	return false

}



