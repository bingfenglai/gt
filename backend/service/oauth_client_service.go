package service

import (
	"errors"

	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/storage"
	"go.uber.org/zap"
)

type oAuthClientServiceImpl struct {
}

func (svc *oAuthClientServiceImpl) GetDetailsByClientId(clientId string) (*entity.Client, error) {

	if clientId == "" {
		return nil, errors.New("client id 不能为空")
	}

	return storage.OAuthClientStorage.SelectOneByClientId(clientId)
}

func (svc *oAuthClientServiceImpl) CheckGrantType(clientId, grantType string) (bool, error) {

	zap.L().Info("检查客户端是否支持该认证方式", zap.String("client_id", clientId), zap.String("grant_type", grantType))
	if clientId == "" || grantType == "" {
		return false, errors.New("当前客户端不允许该认证方式")
	}

	return true, nil
}
