package interfaces

import "github.com/bingfenglai/gt/domain/entity"

// oauth client service
type IOAuthClientService interface {
	GetDetailsByClientId(clientId string) (*entity.Client, error)

	// 检查指定客户端是否允许指定的授权模式
	CheckGrantType(clientId, grantType string) (bool, error)
}
