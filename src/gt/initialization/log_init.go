package initialization

import (
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/handler"
	"github.com/bingfenglai/gt/router"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogConfig() {

	writeSyncer := getLogWriter(config.Conf.Log.Filename, config.Conf.Log.MaxSize, config.Conf.Log.MaxBackups, config.Conf.Log.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(config.Conf.Log.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	global.Log = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(global.Log) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可

	adaptGinLogToZap()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	// 同时输出到控制台跟文件
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	//return zapcore.AddSync(lumberJackLogger)
}

// 将gin日志使用zap输出
func adaptGinLogToZap() {
	router.R.Use(handler.GinZapLogger(), handler.GinZapRecovery(true))

}
