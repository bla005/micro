package service

import (
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

// Config object
type Config struct {
	Service struct {
		LogDir string `yaml:"logdir"`
		Health struct {
			Path string `yaml:"path"`
		} `yaml:"health"`
		Server struct {
			Host    string `yaml:"host"`
			Port    int    `yaml:"port"`
			Ssl     bool   `yaml:"ssl"`
			Timeout struct {
				Read       time.Duration `yaml:"read"`
				Write      time.Duration `yaml:"write"`
				Idle       time.Duration `yaml:"idle"`
				ReadHeader time.Duration `yaml:"read_header"`
			} `yaml:"timeout"`
		} `yaml:"server"`
	} `yaml:"service"`
}

// NewConfig returns an empty config
func NewConfig() *Config {
	return defaultServiceConfig
}

// LoadConfig loads an existing config from the specified path
func LoadConfig(configPath string) (*Config, error) {
	cfg := NewConfig()
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if err := yaml.NewDecoder(file).Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
