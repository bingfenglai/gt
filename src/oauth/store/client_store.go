package store

import (
	"context"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"go.uber.org/zap"
	"strconv"
)

type ClientDbStore struct {
}

func (cs *ClientDbStore) GetByID(_ context.Context, id string) (oauth2.ClientInfo, error) {
	zap.L().Info("客户端id", zap.String("clientId", id))

	client := entity.Client{}

	err := client.GetByBizId(id)

	zap.L().Info("查询到的clientId", zap.String("clientId", client.ClientBizId))

	if err != nil {
		zap.L().Error(err.Error())
	}

	if client.ClientBizId == "" {
		zap.L().Info("无该client")

		return nil, err
	}

	return &models.Client{
		ID:     client.ClientBizId,
		Secret: client.Secret,
		Domain: client.Domain,
		UserID: strconv.Itoa(int(client.CreatedBy)),
	}, err

}
