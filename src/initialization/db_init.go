package initialization

import (
	"encoding/json"
	"os"

	"github.com/bingfenglai/gt/common/constants"
	"github.com/bingfenglai/gt/domain/entity"

	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/global"

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

	if config.Conf.DataBase.InitSchema {
		initSchema()

	}

}

func registerCallback() {
	// err := global.DB.Callback().Create().Register("TenantCallback", callback.CreatedCallback)
	// if err!=nil {
	// 	panic(err)
	// }
	// global.DB.Callback().Create().Remove("gorm:update_time_stamp")
	//_ = global.DB.Callback().Create().Register("gorm:update_time_stamp", CreatedTimeCallback)
	//_ = global.DB.Callback().Update().Replace("gorm:update_time_stamp", UpdatedTimeCallback)
}

func initSchema() {
	_ = global.DB.AutoMigrate(&entity.Role{}, &entity.Dict{}, &entity.DictItem{}, &entity.User{}, &entity.UserRole{},
		&entity.Client{}, &entity.OAuthGrantType{}, &entity.ClientGrantType{},
		&entity.ShortCodeGroup{}, &entity.ShortCode{}, &entity.ShortcodeLog{}, entity.RoleApi{}, &entity.Api{}, &entity.Tenant{})
	if config.Conf.DataBase.InitData {
		initData()
	}
}

func initData() {

	// 插入授权模式
	initGrantTypeData()
}

func initGrantTypeData() {
	password := entity.OAuthGrantType{Status: constants.Normal_Status, Name: "password", Remark: ""}
	password.ID = 1
	authorizationCode := entity.OAuthGrantType{Status: constants.Normal_Status, Name: "authorization_code", Remark: ""}
	authorizationCode.ID = 2
	clientCredentials := entity.OAuthGrantType{Status: constants.Normal_Status, Name: "client_credentials", Remark: ""}
	clientCredentials.ID = 3
	refreshing := entity.OAuthGrantType{Status: constants.Normal_Status, Name: "refresh_token", Remark: ""}
	refreshing.ID = 4
	implicit := entity.OAuthGrantType{Status: constants.Normal_Status, Name: "__implicit", Remark: ""}
	implicit.ID = 5

	grantTypes := make([]*entity.OAuthGrantType, 0)
	grantTypes = append(grantTypes, &password, &authorizationCode, &clientCredentials, &refreshing, &implicit)

	for _, grantType := range grantTypes {
		global.DB.Save(grantType)

	}

}
