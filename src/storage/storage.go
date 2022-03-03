package storage

import "log"

var ShortCodeStorage IShortcodeStorage

var ShortCodeLogStorage IShortcodeLogStorage

var OAuthClientStorage IOAuthClientStorage

func Initstorage() {
	log.Default().Println("初始化 storage")
	ShortCodeStorage = &ShortCodeDbStorage{}

	ShortCodeLogStorage = &ShortCodeLogDbStorage{}

	OAuthClientStorage = &OAuthClientStorageImpl{}

}