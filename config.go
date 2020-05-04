package service

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Service struct {
		Log    bool
		LogDir string
		Server struct {
			Ssl  bool
			Port int
		}
	}
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(configPath string) (*Config, error) {
	cfg := NewConfig()

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	dec := yaml.NewDecoder(file)
	if err := dec.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
