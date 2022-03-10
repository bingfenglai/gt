package config

import (
	"log"
	"os"
	"strconv"
	"strings"

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

	// 热更新
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

	// 配置读取环境变量
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)

	if err != nil {
		log.Default().Println("读取配置信息失败\n", err.Error())
	} else {
		log.Default().Println("config info:\n", Conf.Redis, "\n", Conf.Server, "\nlog: ", Conf.Log,"\nemail",Conf.Email)
	}
	

}

// 配置信息规则检查
func configRoleCheck() {
	if Conf.ShortCode.Size > maxShortCodeGenSize {
		panic("每批次生成的短码数不能大于：" + strconv.Itoa(maxShortCodeGenSize))

	}
}
