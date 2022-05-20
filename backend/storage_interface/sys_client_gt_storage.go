package storage_interface

type IClientGrantTypeStorage interface {
	CheckIsExist(clientId, grantType string) (bool, error)
}
