package storage

import (
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/model/entity"
)

type IOAuthClientStorage interface {

	SelectOneByClientId( clientId string) (*entity.Client,error)

}


type OAuthClientStorageImpl struct {
	
}


func (store *OAuthClientStorageImpl) SelectOneByClientId( clientId string) (*entity.Client,error){

	client := &entity.Client{}
	err :=global.DB.Where(" client_biz_id = ?",clientId).Take(client).Error

	return client,err
}