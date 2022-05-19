package service

import (
	"context"
	"github.com/bingfenglai/gt/global"
)

type baseService struct {
}

func (svc *baseService) Save(ctx context.Context, val interface{}) error {
	return global.DB.WithContext(ctx).Save(val).Error
}

func (svc baseService) SaveBatch(ctx context.Context, val []interface{}) (err error) {

	return
}

func (svc baseService) DeleteById(ctx context.Context, id uint64) (err error) {
	return
}

func (svc baseService) DeleteBatch(ctx context.Context, ids []interface{}) (err error) {

	return
}
