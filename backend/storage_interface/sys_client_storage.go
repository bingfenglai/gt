package storage_interface

import "github.com/bingfenglai/gt/domain/entity"

type IOAuthClientStorage interface {
	SelectOneByClientId(clientId string) (*entity.Client, error)
}
