package initialization

import (
	"encoding/json"
	"log"
	"time"

	"github.com/bingfenglai/gt/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func InitDbConfig() {

	var count = 0

	for {
		db, err := gorm.Open(viper.GetString("database.type"), viper.GetString("database.url"))

		if err != nil {
			log.Default().Println("db 连接错误 ", err.Error())
			time.Sleep(3 * 1e9)
			count++

			if count>3 {
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

	sqlDb.SetMaxOpenConns(viper.GetInt("database.maxOpen"))
	sqlDb.SetMaxIdleConns(viper.GetInt("database.maxConn"))
	dbJson, _ := json.Marshal(sqlDb.Stats())
	log.Default().Println(string(dbJson))

}
