package storage

import "log"

var ShortCodeStorage IShortcodeStorage

var ShortCodeLogStorage IShortcodeLogStorage

var OAuthClientStorage IOAuthClientStorage

var ClientGrantTypeStorage IClientGrantTypeStorage

var UserStorage IUserStorage

var RoleStorage IRoleStorage

var ApiStorage IApiStorage

func Initstorage() {
	log.Default().Println("初始化 storage")
	ShortCodeStorage = &ShortCodeDbStorage{}

	ShortCodeLogStorage = &ShortCodeLogDbStorage{}

	OAuthClientStorage = &OAuthClientStorageImpl{}

	ClientGrantTypeStorage = &ClientGrantTypeStorageImpl{}

	UserStorage = &userDbStorageImpl{}

	RoleStorage = &roleStorage{}

	ApiStorage = &apiStorage{}

}