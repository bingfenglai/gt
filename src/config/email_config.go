package config

// 邮件配置类
type EmailConfig struct {
	// 发送者邮箱
	SenderEmail string
	// 邮件服务器地址
	SmtpServerHost string
	// 密钥
	Auth string `yaml:"auth"`

	Address string

	Enable bool

}
