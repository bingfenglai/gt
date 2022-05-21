package storage_interface

import "context"

type IStorage interface {
	Save(ctx context.Context, val ...interface{}) error

	Delete(ctx context.Context, val interface{}, id ...interface{}) error

	FindOne(ctx context.Context, val interface{}, conds interface{}, fields []string) error
}
