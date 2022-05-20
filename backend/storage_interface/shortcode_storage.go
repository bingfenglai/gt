package storage_interface

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
)

type IShortcodeStorage interface {

	// 跟据md5查找短码
	FindShortcodeByMd5(md5 string) (*entity.ShortCode, error)

	// 跟据短码查找url
	FindOriginalUrlByShortCode(shortcode string) (*entity.ShortCode, error)

	// 存储或者更新
	SaveOrUpdate(ctx context.Context, shortcode *entity.ShortCode) (bool, error)
}
