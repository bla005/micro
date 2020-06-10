package service

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

var startTime time.Time

const (
	dependenciesFile = "dependencies.json"
	endpointsFile    = "endpoints.json"
)

type Dependencies []*Dependency

func (d *Dependencies) ExportJSON() error {
	f, err := os.Create(dependenciesFile)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(d); err != nil {
		return err
	}
	return nil
}

type Endpoints []*Endpoint

func (e *Endpoints) ExportJSON() error {
	f, err := os.Create(endpointsFile)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(e); err != nil {
		return err
	}
	return nil
}

type Service struct {
	Name         string
	Version      string
	config       *Config
	server       *http.Server
	router       *httprouter.Router
	endpoints    Endpoints
	dependencies Dependencies
}

func NewService(name string, version string, config *Config, server *http.Server, router *httprouter.Router) *Service {
	if config == nil {
		panic("nil config")
	}
	if router == nil {
		panic("nil router")
	}
	if server == nil {
		panic("nil server")
	}
	return &Service{
		Name:         name,
		Version:      version,
		config:       config,
		server:       server,
		router:       router,
		endpoints:    make(Endpoints, 0),
		dependencies: make(Dependencies, 0),
	}
}

func (s *Service) Endpoints() []string {
	endpoints := make([]string, len(s.endpoints))
	for _, endpoint := range s.endpoints {
		endpoints = append(endpoints, endpoint.name)
	}

	return endpoints
}

func (s *Service) Dependencies() []string {
	dependencies := make([]string, len(s.dependencies))
	for _, dependency := range s.dependencies {
		dependencies = append(dependencies, dependency.name)
	}
	return dependencies
}

func (s *Service) Uptime() time.Duration {
	return time.Since(startTime)
}

func (s *Service) UseHealthEndpoint() {
	s.router.HandlerFunc("GET", s.config.Service.Health.Path, s.healthHandler)
}

func (s *Service) UseEndpoint(e ...*Endpoint) {
	for _, endpoint := range e {
		s.endpoints = append(s.endpoints, endpoint)
		s.router.HandlerFunc(endpoint.method, endpoint.path, endpoint.handler)
	}
}

func (s *Service) Start() {
	startTime = time.Now()
	if s.config.Service.Server.Ssl {
		// s.server.startWithTls()
	} else {
		start(s.server)
	}
}

func (s *Service) Shutdown() {
	// Is server running?
	if s.Uptime() > 0 {
		shutdown(s.server)
	}
	// Else noop
}
func (s *Service) SetTLSConfig(config *tls.Config) {
	s.server.TLSConfig = config
}

func (s *Service) UseDependency(d ...*Dependency) {
	for _, dependency := range d {
		s.dependencies = append(s.dependencies, dependency)
	}
}

func (s *Service) Health() map[string]string {
	e := make(map[string]string)
	for _, dependency := range s.dependencies {
		if err := dependency.ping(); err != nil {
			e[dependency.name] = "critical"
		}
	}
	return e
}
