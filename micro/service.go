package micro

import (
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	Name    string
	Version float64
	Port    int

	logger *zap.Logger

	server Server

	Endpoints []*Endpoint
}

func NewService(name string, version float64) *Service {
	return &Service{
		Name:      name,
		Version:   version,
		Endpoints: make([]*Endpoint, 1),
	}
}

func (s *Service) GetName() string {
	return s.Name
}
func (s *Service) GetVersion() float64 {
	return s.Version
}
func (s *Service) GetPort() int {
	return s.Port
}
func (s *Service) Start() error {
	router := newRouter(s.Endpoints)
	s.server = newServer(router, "", s.Port)
	if err := s.server.Start(); err != nil {
		return err
	}
	return nil
}

func (s *Service) Stop() {
	s.server.Shutdown()
}

func (s *Service) Health() int {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf(s.server.GetHost(), s.server.GetPort()), 2*time.Second)
	if err != nil {
		return 0
	}
	defer conn.Close()
	return 1
}

func (s *Service) AddEndpoint(endpoint *Endpoint) {
	s.Endpoints = append(s.Endpoints, endpoint)
}
