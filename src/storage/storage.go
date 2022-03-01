package storage

import "log"

var ShortCodeStorage IShortcodeStorage

func Initstorage() {
	log.Default().Println("初始化 storage")
	ShortCodeStorage = &ShortCodeDbStorage{}

}