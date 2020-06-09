package service

import (
	"time"

	"github.com/julienschmidt/httprouter"
)

type Service interface {
	Start()
	Name() string
	Version() string
	Shutdown() error
	Health() map[string]string
	Endpoints() []string
	Dependencies() []string
	Uptime() time.Duration
	UseHealthEndpoint()
	UseEndpoint(e ...*Endpoint)
	UseDependency(d ...*Dependency)
}

type service struct {
	name         string
	version      string
	config       *Config
	server       *Server
	router       *httprouter.Router
	endpoints    Endpoints
	dependencies Dependencies
}

// NewService creates a new service
func NewService(name string, version string, config *Config, server *Server, router *httprouter.Router) (Service, error) {
	if config == nil {
		return nil, ErrNilConfig
	}
	if server == nil {
		return nil, ErrNilServer
	}
	return &service{
		name:         name,
		version:      version,
		config:       config,
		server:       server,
		router:       router,
		endpoints:    make(Endpoints, 0),
		dependencies: make(Dependencies, 0),
	}, nil
}

// Name returns service's name
func (s *service) Name() string {
	return s.name
}

// Version returns service's version
func (s *service) Version() string {
	return s.version
}

// Endpoints returns service's endpoints
func (s *service) Endpoints() []string {
	var endpoints []string
	for i := 0; i < len(s.endpoints); i++ {
		endpoints = append(endpoints, s.endpoints[i].Name())
	}
	return endpoints
}

// Dependencies returns service's dependencies
func (s *service) Dependencies() []string {
	var dependencies []string
	for i := 0; i < len(s.dependencies); i++ {
		dependencies = append(dependencies, s.dependencies[i].Name())
	}
	return dependencies
}

// Uptime returns service's uptime
func (s *service) Uptime() time.Duration {
	if startTime.IsZero() {
		return 0
	}
	return time.Now().Sub(startTime)
}

// UseHealthEndpoint
func (s *service) UseHealthEndpoint() {
	s.router.HandlerFunc("GET", s.config.HealthPath(), s.healthHandler)
}

// AddEndpoints registers services' endpoints
func (s *service) UseEndpoint(e ...*Endpoint) {
	for i := 0; i < len(e); i++ {
		s.endpoints = append(s.endpoints, e[i])
		s.router.HandlerFunc(s.endpoints[i].Method(), s.endpoints[i].Path(), s.endpoints[i].Handler())
	}
	// for i := 0; i < len(s.endpoints); i++ {
	// 	s.router.HandlerFunc(s.endpoints[i].Method(), s.endpoints[i].Path(), s.endpoints[i].Handler())
	// }
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
	if s.Uptime() != 0 {
		s.server.shutdown()
	} else {
		// return ErrServiceNotStarted
		panic("service not started")
	}
	return nil
}

// UseDependency
func (s *service) UseDependency(d ...*Dependency) {
	for i := 0; i < len(d); i++ {
		s.dependencies = append(s.dependencies, d[i])
	}
}

// Health
func (s *service) Health() map[string]string {
	e := make(map[string]string)
	for i := 0; i < len(s.dependencies); i++ {
		if err := s.dependencies[i].Ping(); err != nil {
			e[s.dependencies[i].Name()] = "critical"
		}
	}
	return e
}
