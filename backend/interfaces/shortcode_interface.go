package interfaces

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/domain/params"
)

type IShortCodeService interface {

	// 创建短码并保存
	CreateShortCode(url string, isPerpetual, isMultiplex bool) (*entity.ShortCode, error)

	CreateShortCodeWithContext(params *params.GenShortCodeParams, ctx context.Context) (*entity.ShortCode, error)
	// 根据短码查找原链接
	FindLinkByCode(code string) (*entity.ShortCode, error)

	// 创建临时的短码
	CreatePerpetual(url string) (*entity.ShortCode, error)
}
