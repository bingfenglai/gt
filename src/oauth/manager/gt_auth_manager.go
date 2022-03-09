package manager

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	
	
)

// 增强manager
type CustomOAuthManager struct {
	manager    oauth2.Manager
}

func NewCustomOAuthManager(ma oauth2.Manager) CustomOAuthManager {

	return CustomOAuthManager{
		manager:    ma,
	}
}

func (m *CustomOAuthManager) GenerateAccessToken(ctx context.Context, gt oauth2.GrantType, tgr *oauth2.TokenGenerateRequest) (accessToken oauth2.TokenInfo, err error) {

		// 按照 密码认证模式的配置生成token
		gt = oauth2.PasswordCredentials
		return m.manager.GenerateAccessToken(ctx, gt, tgr)

}

// get the client information
func (m *CustomOAuthManager) GetClient(ctx context.Context, clientID string) (cli oauth2.ClientInfo, err error) {
	return m.manager.GetClient(ctx, clientID)
}

// generate the authorization token(code)
func (m *CustomOAuthManager) GenerateAuthToken(ctx context.Context, rt oauth2.ResponseType, tgr *oauth2.TokenGenerateRequest) (authToken oauth2.TokenInfo, err error) {
	return m.manager.GenerateAuthToken(ctx, rt, tgr)
}

// refreshing an access token
func (m *CustomOAuthManager) RefreshAccessToken(ctx context.Context, tgr *oauth2.TokenGenerateRequest) (accessToken oauth2.TokenInfo, err error) {
	return m.manager.RefreshAccessToken(ctx, tgr)
}

// use the access token to delete the token information
func (m *CustomOAuthManager) RemoveAccessToken(ctx context.Context, access string) (err error) {
	return m.RemoveAccessToken(ctx, access)
}

// use the refresh token to delete the token information
func (m *CustomOAuthManager) RemoveRefreshToken(ctx context.Context, refresh string) (err error) {
	return m.RemoveRefreshToken(ctx, refresh)
}

// according to the access token for corresponding token information
func (m *CustomOAuthManager) LoadAccessToken(ctx context.Context, access string) (ti oauth2.TokenInfo, err error) {
	return m.manager.LoadAccessToken(ctx, access)
}

// according to the refresh token for corresponding token information
func (m *CustomOAuthManager) LoadRefreshToken(ctx context.Context, refresh string) (ti oauth2.TokenInfo, err error) {
	return m.manager.LoadRefreshToken(ctx, refresh)
}