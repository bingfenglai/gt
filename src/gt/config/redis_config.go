package config

type RedisConfig struct {
	Addr      string
	Password  string
	DefaultDb int
	Timeout   int
	PoolSize  int
	MinConn   int
	MaxConn   int
}
