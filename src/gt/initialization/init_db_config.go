package initialization

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"time"
)

var DB *gorm.DB

func InitDbConfig() {

	for {
		db, err := gorm.Open(viper.GetString("database.type"), viper.GetString("database.url"))

		if err != nil {
			log.Default().Println("mysql 连接错误 ", err.Error())
			time.Sleep(1 * 1e9)
			continue
		} else {
			log.Default().Println("数据库连接成功")
			DB = db
			break
		}

	}

	DB.Select("select 1")

	sqlDb := DB.DB()

	sqlDb.SetMaxOpenConns(viper.GetInt("database.maxOpen"))
	sqlDb.SetMaxIdleConns(viper.GetInt("database.maxConn"))
	dbJson, _ := json.Marshal(sqlDb.Stats())
	log.Default().Println(string(dbJson))

}
