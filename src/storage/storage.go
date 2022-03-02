package storage

import "log"

var ShortCodeStorage IShortcodeStorage

var ShortCodeLogStorage IShortcodeLogStorage

func Initstorage() {
	log.Default().Println("初始化 storage")
	ShortCodeStorage = &ShortCodeDbStorage{}

	ShortCodeLogStorage = &ShortCodeLogDbStorage{}

}