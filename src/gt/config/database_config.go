package config

type DataBaseConfig struct {
	DbType  string
	Url     string
	MaxConn int
	MaxOpen int
}
