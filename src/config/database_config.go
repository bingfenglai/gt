package config

type DataBaseConfig struct {
	DbType  string
	Url     string
	MaxConn int
	MaxOpen int
	InitData bool
	InitSchema bool
}
