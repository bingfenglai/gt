package manager

import (
	"context"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/oauth/store"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
)

// 增强manager
type CustomOAuthManager struct {
	originalManager *manage.Manager
}

func NewCustomOAuthManager(ma *manage.Manager) CustomOAuthManager {

	return CustomOAuthManager{
		originalManager: ma,
	}
}

func NewDefaultCustomManager() CustomOAuthManager {
	originalManager := manage.NewDefaultManager()

	// 指定token存储
	originalManager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		DB:       15,
		Password: config.Conf.Redis.Password,
	}))

	// 指定客户端存储
	originalManager.MapClientStorage(&store.ClientDbStore{})

	return CustomOAuthManager{
		originalManager: originalManager,
	}
}

func (m *CustomOAuthManager) GenerateAccessToken(ctx context.Context, gt oauth2.GrantType, tgr *oauth2.TokenGenerateRequest) (accessToken oauth2.TokenInfo, err error) {

	// 按照 密码认证模式的配置生成token
	gt = oauth2.PasswordCredentials
	return m.originalManager.GenerateAccessToken(ctx, gt, tgr)

}

// get the client information
func (m *CustomOAuthManager) GetClient(ctx context.Context, clientID string) (cli oauth2.ClientInfo, err error) {
	return m.originalManager.GetClient(ctx, clientID)
}

// generate the authorization token(code)
func (m *CustomOAuthManager) GenerateAuthToken(ctx context.Context, rt oauth2.ResponseType, tgr *oauth2.TokenGenerateRequest) (authToken oauth2.TokenInfo, err error) {
	return m.originalManager.GenerateAuthToken(ctx, rt, tgr)
}

// refreshing an access token
func (m *CustomOAuthManager) RefreshAccessToken(ctx context.Context, tgr *oauth2.TokenGenerateRequest) (accessToken oauth2.TokenInfo, err error) {
	return m.originalManager.RefreshAccessToken(ctx, tgr)
}

// use the access token to delete the token information
func (m *CustomOAuthManager) RemoveAccessToken(ctx context.Context, access string) (err error) {
	return m.originalManager.RemoveAccessToken(ctx, access)
}

// use the refresh token to delete the token information
func (m *CustomOAuthManager) RemoveRefreshToken(ctx context.Context, refresh string) (err error) {
	return m.originalManager.RemoveRefreshToken(ctx, refresh)
}

// according to the access token for corresponding token information
func (m *CustomOAuthManager) LoadAccessToken(ctx context.Context, access string) (ti oauth2.TokenInfo, err error) {
	return m.originalManager.LoadAccessToken(ctx, access)
}

// according to the refresh token for corresponding token information
func (m *CustomOAuthManager) LoadRefreshToken(ctx context.Context, refresh string) (ti oauth2.TokenInfo, err error) {
	return m.originalManager.LoadRefreshToken(ctx, refresh)
}
