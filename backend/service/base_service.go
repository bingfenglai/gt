package service

import (
	"context"
	"errors"
	"github.com/bingfenglai/gt/global"
)

type baseService struct {
}

func (svc *baseService) Save(ctx context.Context, val interface{}) error {
	if val == nil {
		return errors.New("参数不能为空")
	}

	return global.DB.WithContext(ctx).Save(val).Error
}

func (svc baseService) SaveBatch(ctx context.Context, val []interface{}) (err error) {

	if len(val) == 0 {
		return errors.New("参数不能为空")
	}

	if len(val) == 1 {
		return svc.Save(ctx, val[:1])
	}

	err = global.DB.WithContext(ctx).Save(val).Error

	return
}
