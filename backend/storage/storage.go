package storage

import (
	"github.com/bingfenglai/gt/storage_interface"
	"log"
)

var BaseStorage storage_interface.IStorage

var ShortCodeStorage storage_interface.IShortcodeStorage

var ShortCodeLogStorage storage_interface.IShortcodeLogStorage

var OAuthClientStorage storage_interface.IOAuthClientStorage

var ClientGrantTypeStorage storage_interface.IClientGrantTypeStorage

var UserStorage storage_interface.IUserStorage

var RoleStorage storage_interface.IRoleStorage

var ApiStorage storage_interface.IApiStorage

var TenantStorage storage_interface.ITenantStorage

var SysLogStorage storage_interface.ISysLogStorage

var SysFileStorage storage_interface.ISysFileStorage

func Initstorage() {
	log.Default().Println("初始化 storage")

	BaseStorage = &baseStorage{}

	ShortCodeStorage = &ShortCodeDbStorage{}

	ShortCodeLogStorage = &ShortCodeLogDbStorage{}

	OAuthClientStorage = &OAuthClientStorageImpl{}

	ClientGrantTypeStorage = &ClientGrantTypeStorageImpl{}

	UserStorage = &userDbStorageImpl{}

	RoleStorage = &roleStorage{}

	ApiStorage = &apiStorage{}

	TenantStorage = &tenantStorage{}

	SysLogStorage = &sysLogStorage{}

	SysFileStorage = &SysFileStorageImpl{}

}
