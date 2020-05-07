package service

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type Service struct {
	Name    string
	Version float64

	config *Config

	logger *zap.Logger

	server Server
	router *httprouter.Router

	Endpoints []*Endpoint
	startTime time.Time
}

// Creates a new service
func NewService(name string, version float64, logger *zap.Logger, config *Config) *Service {
	return &Service{
		Name:      name,
		Version:   version,
		config:    config,
		logger:    logger,
		server:    nil,
		router:    nil,
		Endpoints: make([]*Endpoint, 0),
		startTime: time.Now(),
	}
}

// Returns the service name
func (s *Service) GetName() string {
	return s.Name
}

// Returns the service version
func (s *Service) GetVersion() float64 {
	return s.Version
}

// Returns the service port
func (s *Service) GetPort() int {
	return s.config.Service.Server.Port
}

// Returns the endpoints
func (s *Service) GetEndpoints() []*Endpoint {
	return s.Endpoints
}

// Returns the service uptime
func (s *Service) GetUptime() time.Duration {
	return time.Now().Sub(s.startTime)
}

// Initializes the service
func (s *Service) Init() {
	s.router = newRouter()
	s.server = newServer(s.router, s.config)
	if s.config.Service.Log {
		s.logger.Info("Service initialized")
	}
}

// Starts the service
func (s *Service) Start() error {
	if s.config.Service.Server.Ssl {
		if err := s.server.startWithTLS(); err != nil {
			if s.config.Service.Log {
				s.logger.Info("Failed starting service", zap.Error(err))
			}
			return err
		}
	} else {
		if err := s.server.start(); err != nil {
			if s.config.Service.Log {
				s.logger.Info("Failed starting service", zap.Error(err))
			}
			return err
		}
	}
	if s.config.Service.Log {
		s.logger.Info("Service started successfully")
	}
	return nil
}

// Stops the service
func (s *Service) Stop() {
	s.server.shutdown()
	if s.config.Service.Log {
		s.logger.Info("Stopped service successfully")
	}
}

// Add an endpoint to the service
func (s *Service) RegisterEndpoint(handler http.HandlerFunc, method, path string) {
	endpoint := newEndpoint(handler, method, path)
	s.Endpoints = append(s.Endpoints, endpoint)
}

// Register the health endpoint
// which can be used to check the
// health of the service
func (s *Service) RegisterHealthEndpoint() {
	s.router.HandlerFunc("GET", s.config.Service.Health.Endpoint, s.healthHandler)
}
