package service

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type Service interface {
	Init()
	Start() error
	Shutdown()
	GetName() string
	GetVersion() float64
	GetPort() int
	GetRouter() Router
	GetEndpoints() []*Endpoint
	GetUptime() time.Duration
	RegisterEndpoint(handler http.HandlerFunc, method, path string)
	Test(endpoints ...Endpoint)
}

type service struct {
	name      string
	version   float64
	config    *Config
	logger    *zap.Logger
	server    *server
	router    Router
	Endpoints []*Endpoint
}

func (s *service) Test(endpoints ...Endpoint) {
	for i := 0; i < len(endpoints); i++ {
		s.router.HandlerFunc(endpoints[i].Method, endpoints[i].Path, endpoints[i].Handler)
	}
}

// Creates a new service
func NewService(name string, version float64, logger *zap.Logger, config *Config) Service {
	return &service{
		name:      name,
		version:   version,
		config:    config,
		logger:    logger,
		server:    nil,
		router:    nil,
		Endpoints: make([]*Endpoint, 0),
	}
}

// Returns the router
func (s *service) GetRouter() Router {
	return s.router
}

// Returns the service name
func (s *service) GetName() string {
	return s.name
}

// Returns the service version
func (s *service) GetVersion() float64 {
	return s.version
}

// Returns the service port
func (s *service) GetPort() int {
	return s.config.Service.Server.Port
}

// Returns the endpoints
func (s *service) GetEndpoints() []*Endpoint {
	return s.Endpoints
}

// Returns the service uptime
func (s *service) GetUptime() time.Duration {
	return time.Now().Sub(startTime)
}

// Initializes the service
func (s *service) Init() {
	s.router = newRouter()
	s.server = newServer(s.router, s.config)
	s.router.HandlerFunc("GET", s.config.Service.Health.Path, s.healthHandler)
	// s.logger.Info("Service initialized")
}

// Starts the service
func (s *service) Start() error {
	if s.config.Service.Server.Ssl {
		if err := s.server.startWithTLS(); err != nil {
			// s.logger.Info("Failed starting service", zap.Error(err))
			return err
		}
	} else {
		if err := s.server.start(); err != nil {
			// s.logger.Info("Failed starting service", zap.Error(err))
			return err
		}
	}
	// s.logger.Info("Service started successfully")
	return nil
}

// Stops the service
func (s *service) Shutdown() {
	s.server.shutdown()
	// s.logger.Info("Stopped service successfully")
}

// Add an endpoint to the service
func (s *service) RegisterEndpoint(handler http.HandlerFunc, method, path string) {
	endpoint := newEndpoint(handler, method, path)
	s.Endpoints = append(s.Endpoints, endpoint)
}
