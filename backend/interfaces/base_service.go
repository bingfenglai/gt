package interfaces

import "context"

type Service interface {
	Save(ctx context.Context, val interface{}) error

	SaveBatch(ctx context.Context, val []interface{}) error

	DeleteById(ctx context.Context, id uint64) error

	DeleteBatch(ctx context.Context, ids []interface{}) error
}
