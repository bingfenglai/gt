package test

import (
	"log"
	"testing"

	"github.com/bingfenglai/gt/global"
)

func TestDb(t *testing.T) {
	r := 0
	global.DB.Raw("SELECT 1 FROM tb_sys_client_grant_type AS t1 LEFT JOIN tb_sys_grant_type AS t2 ON t1.grant_type_id = t2.id WHERE t1.client_biz_id = ? AND t2.name = ? LIMIT 1",999,"password").Scan(&r)

	log.Default().Println("结果： ",r)	
}