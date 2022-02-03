package initialization

import (
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"time"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/model/entity"
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

	registerCallback()

	initSchema()

}

func registerCallback() {
	global.DB.Callback().Create().Replace("gorm:update_time_stamp", CreatedTimeCallback)
	global.DB.Callback().Update().Replace("gorm:update_time_stamp", UpdatedTimeCallback)
}

func initSchema() {
	global.DB.AutoMigrate(&entity.ShortCodeGroup{}, &entity.Role{}, &entity.Dict{}, &entity.User{})

	lg := entity.ShortCodeGroup{
		GroupName: "default",
		CreatedBy: 0,
	}

	global.DB.Begin().Save(&lg).Commit()
}

func CreatedTimeCallback(scope *gorm.Scope) {

	if scope.HasError() {
		return
	}

	field, ok := scope.FieldByName("createdAt")

	if ok {
		if !field.HasDefaultValue {

			zap.L().Info("插入创建时间")
			_ = field.Set(time.Now().Local().Format("2006-01-02 15:04:05"))

		}
	}
}

func UpdatedTimeCallback(scope *gorm.Scope) {

	if scope.HasError() {
		return
	}

	field, ok := scope.FieldByName("updatedAt")

	if ok {
		if !field.HasDefaultValue {

			_ = field.Set(time.Now().Local().Format("2006-01-02 15:04:05"))

		}
	}
}
