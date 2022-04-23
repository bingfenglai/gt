package test

import (
	"context"
	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
	"testing"
)

func TestGenShortCode(t *testing.T) {

	codeParams := params.GenShortCodeParams{
		OriginalLink: "https://www.baidu.com",
		IsPerpetual:  true,
	}
	if sc, err := service.ShortCodeService.CreateShortCodeWithContext(&codeParams, context.Background()); err != nil {
		t.Error(err)
	} else {
		zap.L().Info("生成的短码", zap.Any("sc", sc))
	}

}
