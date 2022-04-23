package storage

import (
	"context"
	"errors"

	"github.com/bingfenglai/gt/common/constants"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
)

type IShortcodeStorage interface {

	// 跟据md5查找短码
	FindShortcodeByMd5(md5 string) (*entity.ShortCode, error)

	// 跟据短码查找url
	FindOriginalUrlByShortCode(shortcode string) (*entity.ShortCode, error)

	// 存储或者更新
	SaveOrUpdate(ctx context.Context, shortcode *entity.ShortCode) (bool, error)
}

type ShortCodeDbStorage struct {
}

func (store *ShortCodeDbStorage) FindShortcodeByMd5(md5 string) (*entity.ShortCode, error) {
	sc := entity.ShortCode{}

	if err := global.DB.Where("md5 = ? AND `status` = ?", md5, constants.Normal_Status).Take(&sc).Error; err != nil {
		return nil, err
	}

	return &sc, nil

}

func (store *ShortCodeDbStorage) FindOriginalUrlByShortCode(code string) (*entity.ShortCode, error) {
	shortcode := entity.ShortCode{}

	if err := global.DB.Select("id", "original").Where("short_code = ? AND `status` = ? ", code, constants.Normal_Status).Take(&shortcode).Error; err != nil {
		return nil, err
	}

	return &shortcode, nil
}

func (store *ShortCodeDbStorage) SaveOrUpdate(ctx context.Context, shortcode *entity.ShortCode) (bool, error) {

	if shortcode == nil {
		return false, errors.New("短码入参不能为空")
	}

	if shortcode.ID == 0 {

		if err := global.DB.WithContext(ctx).Create(shortcode).Error; err != nil {
			return false, err
		} else {
			return true, nil
		}
	}

	err := global.DB.Save(shortcode).Error

	return err == nil, err
}
