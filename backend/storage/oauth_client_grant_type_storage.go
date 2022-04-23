package storage

import (
	"github.com/bingfenglai/gt/global"
)

type IClientGrantTypeStorage interface {
	CheckIsExist(clientId, grantType string) (bool, error)
}

type ClientGrantTypeStorageImpl struct {
}

func (store *ClientGrantTypeStorageImpl) CheckIsExist(clientId, grantType string) (bool, error) {

	r := 0
	
	err := global.DB.Raw("SELECT 1 FROM tb_sys_client_grant_type AS t1 LEFT JOIN tb_sys_grant_type AS t2 ON t1.grant_type_id = t2.id WHERE t1.client_biz_id = ? AND t2.name = ? LIMIT 1",999,"password").Scan(&r).Error



	return r==1,err

}
