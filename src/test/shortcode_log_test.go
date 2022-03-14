package test

import (
	"testing"

	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/storage"
	"go.uber.org/zap"
)

func TestLogSave(t *testing.T) {

	log := entity.ShortcodeLog{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 Edg/98.0.1108.62",
	}

	log.ID = 1

	if flag, err := storage.ShortCodeLogStorage.SaveOrUpdate(&log); !flag {
		zap.L().Error(err.Error())
		t.Fail()
	} else {
		zap.L().Info("true")
	}

}
