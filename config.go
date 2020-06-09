package service

import (
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

//
// type Config interface {
// 	LogDir() string
// 	HealthPath() string
// 	ServerHost() string
// 	ServerPort() int
// 	ServerSsl() bool
// 	ServerReadTimeout() time.Duration
// 	ServerWriteTimeout() time.Duration
// 	ServerIdleTimeout() time.Duration
// 	ServerReadHeaderTimeout() time.Duration
// 	SetLogDir(dir string)
// 	SetHealthPath(path string)
// 	SetHost(host string)
// 	SetPort(port int)
// 	SetSsl(ssl bool)
// 	SetReadTimeout(timeout time.Duration)
// 	SetWriteTimeout(timeout time.Duration)
// 	SetIdleTimeout(timeout time.Duration)
// 	SetReadHeaderTimeout(timeout time.Duration)
// }
//

// Config structure
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

// NewConfig creates a Config with default settings
func NewConfig() *Config {
	return defaultServiceConfig
}
func (c *Config) LogDir() string {
	return c.Service.LogDir
}
func (c *Config) HealthPath() string {
	return c.Service.Health.Path
}
func (c *Config) ServerHost() string {
	return c.Service.Server.Host
}
func (c *Config) ServerPort() int {
	return c.Service.Server.Port
}
func (c *Config) ServerSsl() bool {
	return c.Service.Server.Ssl
}
func (c *Config) ServerReadTimeout() time.Duration {
	return c.Service.Server.Timeout.Read
}
func (c *Config) ServerWriteTimeout() time.Duration {
	return c.Service.Server.Timeout.Write
}
func (c *Config) ServerIdleTimeout() time.Duration {
	return c.Service.Server.Timeout.Idle
}
func (c *Config) ServerReadHeaderTimeout() time.Duration {
	return c.Service.Server.Timeout.ReadHeader
}

func (c *Config) SetLogDir(dir string) {
	c.Service.LogDir = dir
}

func (c *Config) SetHealthPath(path string) {
	c.Service.Health.Path = path
}
func (c *Config) SetHost(host string) {
	c.Service.Server.Host = host
}
func (c *Config) SetPort(port int) {
	c.Service.Server.Port = port
}
func (c *Config) SetSsl(ssl bool) {
	c.Service.Server.Ssl = ssl
}
func (c *Config) SetReadTimeout(timeout time.Duration) {
	c.Service.Server.Timeout.Read = timeout
}
func (c *Config) SetWriteTimeout(timeout time.Duration) {
	c.Service.Server.Timeout.Write = timeout
}
func (c *Config) SetIdleTimeout(timeout time.Duration) {
	c.Service.Server.Timeout.Idle = timeout
}
func (c *Config) SetReadHeaderTimeout(timeout time.Duration) {
	c.Service.Server.Timeout.ReadHeader = timeout
}

// LoadConfig loads an existing Config from the specified path
func LoadConfig(path string) (*Config, error) {
	c := NewConfig()
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if err := yaml.NewDecoder(file).Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}
