package service

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Service struct {
		Log        bool
		LogsFolder string
		Server     struct {
			UseSsl   bool
			HttpPort int
		}
	}
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(configPath string) (*Config, error) {
	cfg := NewConfig()

	file, err := os.OpenFile(configPath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	dec := yaml.NewDecoder(file)
	if err := dec.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
