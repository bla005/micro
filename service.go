package service

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// Service object
type Service interface {
	Init()
	Start() error
	Shutdown()
	GetName() string
	GetVersion() string
	GetPort() int
	GetRouter() *httprouter.Router
	GetEndpoints() []string
	GetUptime() time.Duration
	RegisterEndpoint(name string, handler http.HandlerFunc, method, path string)
	RegisterEndpoints()
	RegisterHealthEndpoint()
	SetTLSConfig(config *tls.Config)
	GetTLSConfig() *tls.Config
}

type service struct {
	name      string
	version   string
	config    *Config
	logger    *zap.Logger
	server    *server
	router    *httprouter.Router
	endpoints []*Endpoint
}

// NewService creates a new service
func NewService(name string, version string, logger *zap.Logger, config *Config) (Service, error) {
	if logger == nil {
		return nil, ErrNilLogger
	}
	if config == nil {
		return nil, ErrNilConfig
	}
	return &service{
		name:      name,
		version:   version,
		config:    config,
		logger:    logger,
		server:    nil,
		router:    nil,
		endpoints: []*Endpoint{},
	}, nil
}

// GetRouter returns service's router
func (s *service) GetRouter() *httprouter.Router {
	return s.router
}

// GetName returns service's name
func (s *service) GetName() string {
	return s.name
}

// GetVersion returns service's version
func (s *service) GetVersion() string {
	return s.version
}

// GetPort returns service's port
func (s *service) GetPort() int {
	return s.config.Service.Server.Port
}

// GetEndpoints returns service's endpoints
func (s *service) GetEndpoints() []string {
	var endpoints []string
	for i := 0; i < len(s.endpoints); i++ {
		endpoints = append(endpoints, s.endpoints[i].Name)
	}
	return endpoints
}

// GetUptime returns service's uptime
func (s *service) GetUptime() time.Duration {
	return time.Now().Sub(startTime)
}

// Init initializes the router and the server
func (s *service) Init() {
	s.router = newRouter()
	s.server = newServer(s.router, s.config)
}

func (s *service) RegisterHealthEndpoint() {
	s.router.HandlerFunc("GET", s.config.Service.Health.Path, s.healthHandler)
}

// SetTLSConfig makes the server use a specific TLS config
func (s *service) SetTLSConfig(config *tls.Config) {
	s.server.setTLSConfig(config)
}

// GetTLSConfig returns the TLS config in use
func (s *service) GetTLSConfig() *tls.Config {
	return s.server.getTLSConfig()
}

// Start starts the service
func (s *service) Start() error {
	if s.config.Service.Server.Ssl {
		if err := s.server.startWithTLS(); err != nil {
			return err
		}
	} else {
		s.server.start()
		/*
			if err := s.server.start(); err != nil {
				return err
			}
		*/
	}
	startTime = time.Now()
	return nil
}

// Shudtown stops the service
func (s *service) Shutdown() {
	s.server.shutdown()
}
