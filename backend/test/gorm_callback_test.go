package test

import (
	"testing"

	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"
)

func TestGormCallBack(t *testing.T) {
	te := entity.Tenant{}
	global.DB.First(&te)
	
	global.DB.Model(&te).Update("name","海南省省厅")
}
