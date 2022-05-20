package service_interfaces

import "context"

type Service interface {
	Save(ctx context.Context, val interface{}) error

	SaveBatch(ctx context.Context, val []interface{}) error
}
