package config

import (
	"log"
	"os"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	Redis     RedisConfig
	Server    ServerConfig
	DataBase  DataBaseConfig
	Swagger   SwaggerConfig
	Log       LogConfig
	Cache     CacheConfig
	ShortCode ShortCodeConfig
	Captcha   CaptchaConfig
	Auth      OAuth2Config
	Encrypt EncryptConfig
	Email EmailConfig
}

func init() {
	log.Default().Println("装载配置文件信息")
	LoadConfig()
	//zap.L().Info("装载配置文件信息")
	configRoleCheck()

	viper.OnConfigChange(func(in fsnotify.Event) {
		LoadConfig()
		//zap.L().Info("装载配置文件信息")
		configRoleCheck()
	})

}

// 加载配置文件

func LoadConfig() {
	// 获取当前工作目录
	workDir, _ := os.Getwd()
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)

	if err != nil {
		log.Default().Println("读取配置信息失败\n", err.Error())
	} else {
		log.Default().Println("config info:\n", Conf.Redis, "\n", Conf.Server, "\nlog: ", Conf.Log)
	}

}

// 配置信息规则检查
func configRoleCheck() {
	if Conf.ShortCode.Size > maxShortCodeGenSize {
		panic("每批次生成的短码数不能大于：" + strconv.Itoa(maxShortCodeGenSize))

	}
}
