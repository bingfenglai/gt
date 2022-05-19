package service

import (
	"context"
	"github.com/bingfenglai/gt/global"
)

type BaseService struct {
}

func (svc *BaseService) Save(ctx context.Context, val interface{}) error {
	return global.DB.WithContext(ctx).Save(val).Error
}

func (svc BaseService) SaveBatch(ctx context.Context, val []interface{}) (err error) {

	return
}

func (svc BaseService) DeleteById(ctx context.Context, id uint64) (err error) {
	return
}

func (svc BaseService) DeleteBatch(ctx context.Context, ids []interface{}) (err error) {

	return
}
