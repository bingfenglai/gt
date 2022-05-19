package interfaces

import "github.com/bingfenglai/gt/domain/entity"

type IShortCodeLogService interface {
	// 保存
	Save(shortcodeLog *entity.ShortcodeLog) (bool, error)

	// 创建
	Create(shorCodeId uint64, userAgent string, ip string) (bool, error)
}
