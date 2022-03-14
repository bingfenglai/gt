package storage

import (
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
)

type IShortcodeLogStorage interface {
	SaveOrUpdate(shortcodeLog *entity.ShortcodeLog) (bool, error)
}

type ShortCodeLogDbStorage struct {
}

func (store *ShortCodeLogDbStorage) SaveOrUpdate(shortcodeLog *entity.ShortcodeLog) (bool, error) {
	old := entity.ShortcodeLog{}
	if shortcodeLog.ID == 0 {
		if err := global.DB.Create(shortcodeLog).Error; err != nil {
			return false, err
		}

		return true, nil
	}

	global.DB.Where(" id = ?", shortcodeLog.ID).First(&old)

	if old.ID == 0 {
		err := global.DB.Create(shortcodeLog).Error

		return err == nil, err
	}

	shortcodeLog.ID = old.ID

	err := global.DB.Updates(shortcodeLog).Error

	return err == nil, err

}
