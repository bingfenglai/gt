package storage_interface

import "github.com/bingfenglai/gt/domain/entity"

type IShortcodeLogStorage interface {
	SaveOrUpdate(shortcodeLog *entity.ShortcodeLog) (bool, error)
}
