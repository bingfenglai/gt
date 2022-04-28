package session
var UserSessionService IUserSessionService

func init()  {
	UserSessionService = &userSessionServiceImpl{}
}