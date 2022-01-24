package initialization

import (
	"encoding/json"
	"github.com/bingfenglai/gt/config"
	"log"
	"time"

	"github.com/bingfenglai/gt/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func initDbConfig() {

	var count = 0

	for {

		db, err := gorm.Open(config.Conf.DataBase.DbType, config.Conf.DataBase.Url)

		if err != nil {
			log.Default().Println("db 连接错误 ", err.Error())
			time.Sleep(3 * 1e9)
			count++

			if count > 3 {
				panic("连接数据库失败")
			}

			continue
		} else {
			log.Default().Println("数据库连接成功")
			global.DB = db
			break
		}

	}

	global.DB.Select("select 1")

	sqlDb := global.DB.DB()
	sqlDb.SetMaxOpenConns(config.Conf.DataBase.MaxOpen)
	sqlDb.SetMaxIdleConns(config.Conf.DataBase.MaxConn)
	dbJson, _ := json.Marshal(sqlDb.Stats())
	log.Default().Println(string(dbJson))

}
