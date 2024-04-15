package configs

import (
	"github.com/pelletier/go-toml/v2"
	"io"
	"os"
)

type Config struct {
	DB struct {
		Host         string `toml:"host"`
		User         string `toml:"user"`
		Password     string `toml:"password"`
		DatabaseName string `toml:"database_name"`
		Port         int    `toml:"port"`
	} `toml:"db"`
}

func ReadConf() (*Config, error) {
	file, err := os.Open("internal/configs/config.toml")
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err = toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, err
}
