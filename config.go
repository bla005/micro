package service

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Service struct {
		Log         bool   `yaml:"log"`
		LogDir      string `yaml:"logdir"`
		HealthCheck bool   `yaml:"health_check"`
		Health      struct {
			Check    bool   `yaml:"check"`
			Timeout  uint64 `yaml:"timeout"`
			Endpoint string `yaml:"endpoint"`
		} `yaml:"health"`
		Server struct {
			Ssl  bool `yaml:"ssl"`
			Port int  `yaml:"port"`
		} `yaml:"server"`
	} `yaml:"service"`
}

func newConfig() *Config {
	return &Config{}
}

// Loads config from an existing config yaml file
func LoadConfig(configPath string) (*Config, error) {
	cfg := newConfig()

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
