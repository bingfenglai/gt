package config


var ServerConfigInfo ServerConfig

type ServerConfig struct {
	Address           string
	Port int
	ActiveProfiles string
}
