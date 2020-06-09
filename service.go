package service

import (
	"time"

	"github.com/julienschmidt/httprouter"
)

// Service structure
type Service interface {
	Name() string
	Version() string
	Start()
	Shutdown() error
	Health() map[string]string
	Endpoints() []string
	Dependencies() []string
	Uptime() time.Duration
	RegisterEndpoints()
	RegisterHealthEndpoint()
}

type service struct {
	name         string
	version      string
	config       *Config
	server       *Server
	router       *httprouter.Router
	endpoints    []*Endpoint
	dependencies []*Dependency
}

// NewService creates a new service
func NewService(name string, version string, config *Config, server *Server, router *httprouter.Router) (Service, error) {
	if config == nil {
		return nil, ErrNilConfig
	}
	return &service{
		name:         name,
		version:      version,
		config:       config,
		server:       server,
		router:       router,
		endpoints:    make([]*Endpoint, 0),
		dependencies: make([]*Dependency, 0),
	}, nil
}

// GetName returns service's name
func (s *service) Name() string {
	return s.name
}

// GetVersion returns service's version
func (s *service) Version() string {
	return s.version
}

// GetEndpoints returns service's endpoints
func (s *service) Endpoints() []string {
	var endpoints []string
	for i := 0; i < len(s.endpoints); i++ {
		endpoints = append(endpoints, s.endpoints[i].Name)
	}
	return endpoints
}

// GetDependencies returns service's dependencies
func (s *service) Dependencies() []string {
	var dependencies []string
	for i := 0; i < len(s.dependencies); i++ {
		dependencies = append(dependencies, s.dependencies[i].Name)
	}
	return dependencies
}

// GetUptime returns service's uptime
func (s *service) Uptime() time.Duration {
	if startTime.IsZero() {
		return 0
	}
	return time.Now().Sub(startTime)
}

func (s *service) RegisterHealthEndpoint() {
	s.router.HandlerFunc("GET", s.config.HealthPath(), s.healthHandler)
}

// Start starts the service
func (s *service) Start() {
	startTime = time.Now()
	if s.config.ServerSsl() {
		// s.server.startWithTls()
	} else {
		s.server.start()
	}
}

// Shutdown stops the service
func (s *service) Shutdown() error {
	if s.Uptime() == 0 {
		s.server.shutdown()
	} else {
		// return ErrServiceNotStarted
		panic("service not started")
	}
	return nil
}
