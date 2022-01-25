package config

// 日志配置
type LogConfig struct {
	// 日志写入文件的路径
	Filename string
	// 当前日志记录级别
	Level string
	// 日志编码方式
	Encoder string
	// 日志文件切割阈值
	MaxSize int
	// 文件保存天数
	MaxAge int

	//保留的旧日志文件做大个数
	MaxBackups int
}
