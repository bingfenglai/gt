package manager

import "github.com/go-oauth2/oauth2/v4"

// MapClientStorage mapping the client store interface

func (m *CustomOAuthManager) MapClientStorage(store oauth2.ClientStore) {
	m.originalManager.MapClientStorage(store)
}


// MapTokenStorage mapping the token store interface
func (m *CustomOAuthManager) MapTokenStorage(store oauth2.TokenStore) {
	m.originalManager.MapTokenStorage(store)
}