package service

import (
	"bytes"

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

func LoadConfig() (*Config, error) {
	cfg := NewConfig()
	plm := yaml.ReferenceDirs(".")
	buf := bytes.NewBufferString("")
	dec := yaml.NewDecoder(buf, plm)
	if err := dec.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
