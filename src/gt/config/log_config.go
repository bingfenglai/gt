package config

// 日志配置
type LogConfig struct {
	// 日志写入文件的路径
	Path string
	// 当前日志记录级别
	Level string
	// 日志编码方式
	Encoder string
	// 日志文件切割阈值
	MaxSize int
	// 文件保存天数
	MaxAge int

	//备份数
	maxBackup int
}
