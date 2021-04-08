package api

import "StandartWebServer/storage"

type Config struct {
	//Port
	BindAddr string `toml:"bind_addr"`

	//Logger level
	LoggerLevel string `toml:"logger_level"`

	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage: storage.NewConfig(),
	}
}
