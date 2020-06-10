package service

import (
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

// DefaultConfig is the default service config
var DefaultConfig = &Config{
	Service: &serviceConfig{
		LogDir: "logs",
		Health: &healthConfig{
			Path: "/health",
		},
		Server: &serverConfig{
			Host: "",
			Port: 8088,
			Timeout: &timeoutConfig{
				Read:       time.Second * 2,
				Write:      time.Second * 2,
				Idle:       time.Second * 20,
				ReadHeader: time.Second * 5,
			},
		},
	},
}

type Config struct {
	Service *serviceConfig `yaml:"service"`
}

type serviceConfig struct {
	LogDir string        `yaml:"logdir"`
	Health *healthConfig `yaml:"health"`
	Server *serverConfig `yaml:"server"`
}

type healthConfig struct {
	Path string `yaml:"path"`
}

type serverConfig struct {
	Host    string         `yaml:"host"`
	Port    int            `yaml:"port"`
	Ssl     bool           `yaml:"ssl"`
	Timeout *timeoutConfig `yaml:"timeout"`
}

type timeoutConfig struct {
	Read       time.Duration `yaml:"read"`
	Write      time.Duration `yaml:"write"`
	Idle       time.Duration `yaml:"idle"`
	ReadHeader time.Duration `yaml:"read_header"`
}

// NewConfig returns an empty config
func NewConfig() *Config {
	return &Config{}
}

// LoadConfig loads a config from a path
func OpenConfig(file string) (*Config, error) {
	cfg := NewConfig()
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) Save(file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := yaml.NewEncoder(f).Encode(c); err != nil {
		return err
	}
	return nil
}
