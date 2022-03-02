package initialization

import (
	"encoding/json"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/model/entity"
	_ "github.com/go-sql-driver/mysql"
)

func initDbConfig() {

	var count = 0

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Microsecond, // 慢 SQL 阈值
			LogLevel:                  logger.Warn,      // 日志级别
			IgnoreRecordNotFoundError: true,             // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,            // 禁用彩色打印
		},
	)

	for {

		//db, err := gorm.Open(config.Conf.DataBase.DbType, config.Conf.DataBase.Url)
		db, err := gorm.Open(mysql.Open(config.Conf.DataBase.Url), &gorm.Config{
			Logger: newLogger,
			},
		)
		//db.SetLogger(logger.Default.LogMode(logger.Warn))

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

	sqlDb, _ := global.DB.DB()

	sqlDb.SetMaxOpenConns(config.Conf.DataBase.MaxOpen)
	sqlDb.SetMaxIdleConns(config.Conf.DataBase.MaxConn)

	dbJson, _ := json.Marshal(sqlDb.Stats())
	log.Default().Println(string(dbJson))

	registerCallback()

	initSchema()

}

func registerCallback() {
	global.DB.Callback().Create().Remove("gorm:update_time_stamp")
	//_ = global.DB.Callback().Create().Register("gorm:update_time_stamp", CreatedTimeCallback)
	//_ = global.DB.Callback().Update().Replace("gorm:update_time_stamp", UpdatedTimeCallback)
}

func initSchema() {
	_ = global.DB.AutoMigrate(&entity.ShortCodeGroup{}, &entity.Role{}, &entity.Dict{}, &entity.User{},
		&entity.Client{},&entity.ShortCode{},&entity.ShortcodeLog{})

	// lg := entity.ShortCodeGroup{
	// 	GroupName: "default",
	// 	CreatedBy: 0,
	// }

	// global.DB.Begin().Save(&lg).Commit()
}

func CreatedTimeCallback(db *gorm.DB) {

	field := db.Statement.Schema.LookUpField("CreatedAt")

	if field != nil {
		if !field.HasDefaultValue {

			_ = field.Set(db.Statement.ReflectValue, time.Now().Local().Format("2006-01-02 15:04:05"))

		}
	}
}

func UpdatedTimeCallback(db *gorm.DB) {

	field := db.Statement.Schema.LookUpField("UpdatedAt")

	if field != nil {
		if !field.HasDefaultValue {

			_ = field.Set(db.Statement.ReflectValue, time.Now().Local().Format("2006-01-02 15:04:05"))

		}
	}
}
