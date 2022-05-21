package service

import (
	"context"
	"errors"
	"github.com/bingfenglai/gt/storage"
	"go.uber.org/zap"
)

type baseService struct {
}

func (svc *baseService) Save(ctx context.Context, val interface{}) error {
	if val == nil {
		return errors.New("参数不能为空")
	}

	return storage.BaseStorage.Save(ctx, val)
}

func (svc *baseService) FindOne(ctx context.Context, val interface{}, conds interface{}, fields []string) error {

	if val == nil {
		return errors.New("参数val不能为空")
	}

	if fields == nil || len(fields) == 0 {
		zap.L().Warn("建议声明要查询的具体字段，避免使用*通配符")
	}

	return storage.BaseStorage.FindOne(ctx, val, conds, fields)

}

func (svc *baseService) Delete(ctx context.Context, val interface{}, id ...interface{}) error {

	if val == nil {
		return errors.New("请指定类型")
	}

	return storage.BaseStorage.Delete(ctx, val, id...)

}
