package service_interfaces

import "context"

type Service interface {
	Save(ctx context.Context, val interface{}) error

	FindOne(ctx context.Context, val interface{}, conds interface{}, fields []string) error

	Delete(ctx context.Context, val interface{}, id ...interface{}) error
}
