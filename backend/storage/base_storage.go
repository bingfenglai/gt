package storage

import (
	"context"
	"errors"
	"github.com/bingfenglai/gt/global"
)

type baseStorage struct {
}

func (store *baseStorage) Save(ctx context.Context, vals ...interface{}) (err error) {

	if vals == nil {
		return errors.New("参数不能为空")
	}

	for _, val := range vals {
		err = global.DB.WithContext(ctx).Save(val).Error
	}

	return

}

func (store *baseStorage) Delete(ctx context.Context, val interface{}, conds ...interface{}) (err error) {

	for _, cond := range conds {
		err = global.DB.WithContext(ctx).Delete(val, cond).Error
	}

	return
}

func (store *baseStorage) Find(ctx context.Context, val interface{}, conds interface{}) (err error) {
	global.DB.WithContext(ctx).First(val, conds)
	return
}
