package server

import (
	"github.com/go-oauth2/oauth2/v4"
	svr "github.com/go-oauth2/oauth2/v4/server"
	
)

// SetTokenType token type
func (s *CustomOAuthServer) SetTokenType(tokenType string) {
	s.originalServer.Config.TokenType = tokenType
}

// SetAllowGetAccessRequest to allow GET requests for the token
func (s *CustomOAuthServer) SetAllowGetAccessRequest(allow bool) {
	s.originalServer.Config.AllowGetAccessRequest = allow
}

// SetAllowedResponseType allow the authorization types
func (s *CustomOAuthServer) SetAllowedResponseType(types ...oauth2.ResponseType) {
	s.originalServer.Config.AllowedResponseTypes = types
}

// SetAllowedGrantType allow the grant types
func (s *CustomOAuthServer) SetAllowedGrantType(types ...oauth2.GrantType) {
	s.originalServer.Config.AllowedGrantTypes = types
}

// SetClientInfoHandler get client info from request
func (s *CustomOAuthServer) SetClientInfoHandler(handler svr.ClientInfoHandler) {
	s.originalServer.ClientInfoHandler = handler
}

// SetClientAuthorizedHandler check the client allows to use this authorization grant type
func (s *CustomOAuthServer) SetClientAuthorizedHandler(handler svr.ClientAuthorizedHandler) {
	s.originalServer.ClientAuthorizedHandler = handler
}

// SetClientScopeHandler check the client allows to use scope
func (s *CustomOAuthServer) SetClientScopeHandler(handler svr.ClientScopeHandler) {
	s.originalServer.ClientScopeHandler = handler
}

// SetUserAuthorizationHandler get user id from request authorization
func (s *CustomOAuthServer) SetUserAuthorizationHandler(handler svr.UserAuthorizationHandler) {
	s.originalServer.UserAuthorizationHandler = handler
}

// SetPasswordAuthorizationHandler get user id from username and password
func (s *CustomOAuthServer) SetPasswordAuthorizationHandler(handler svr.PasswordAuthorizationHandler) {
	s.originalServer.PasswordAuthorizationHandler = handler
}

// SetRefreshingScopeHandler check the scope of the refreshing token
func (s *CustomOAuthServer) SetRefreshingScopeHandler(handler svr.RefreshingScopeHandler) {
	s.originalServer.RefreshingScopeHandler = handler
}

// SetRefreshingValidationHandler check if refresh_token is still valid. eg no revocation or other
func (s *CustomOAuthServer) SetRefreshingValidationHandler(handler svr.RefreshingValidationHandler) {
	s.originalServer.RefreshingValidationHandler = handler
}

// SetResponseErrorHandler response error handling
func (s *CustomOAuthServer) SetResponseErrorHandler(handler svr.ResponseErrorHandler) {
	s.originalServer.ResponseErrorHandler = handler
}

// SetInternalErrorHandler internal error handling
func (s *CustomOAuthServer) SetInternalErrorHandler(handler svr.InternalErrorHandler) {
	s.originalServer.InternalErrorHandler = handler
}

// SetPreRedirectErrorHandler sets the PreRedirectErrorHandler in current CustomOAuthServer instance
func (s *CustomOAuthServer) SetPreRedirectErrorHandler(handler svr.PreRedirectErrorHandler) {
	s.originalServer.PreRedirectErrorHandler = handler
}

// SetExtensionFieldsHandler in response to the access token with the extension of the field
func (s *CustomOAuthServer) SetExtensionFieldsHandler(handler svr.ExtensionFieldsHandler) {
	s.originalServer.ExtensionFieldsHandler = handler
}

// SetAccessTokenExpHandler set expiration date for the access token
func (s *CustomOAuthServer) SetAccessTokenExpHandler(handler svr.AccessTokenExpHandler) {
	s.originalServer.AccessTokenExpHandler = handler
}

// SetAuthorizeScopeHandler set scope for the access token
func (s *CustomOAuthServer) SetAuthorizeScopeHandler(handler svr.AuthorizeScopeHandler) {
	s.originalServer.AuthorizeScopeHandler = handler
}

// SetResponseTokenHandler response token handing
func (s *CustomOAuthServer) SetResponseTokenHandler(handler svr.ResponseTokenHandler) {
	s.originalServer.ResponseTokenHandler = handler
}