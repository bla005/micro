package service

import (
	"fmt"
	"net"
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
}

func NewService(name string, version float64, logger *zap.Logger, config *Config) *Service {
	return &Service{
		Name:      name,
		Version:   version,
		config:    config,
		logger:    logger,
		server:    nil,
		router:    nil,
		Endpoints: make([]*Endpoint, 0),
	}
}

func (s *Service) GetName() string {
	return s.Name
}
func (s *Service) GetVersion() float64 {
	return s.Version
}
func (s *Service) GetPort() int {
	return s.config.Service.Server.HttpPort
}
func (s *Service) InitRouter() {
	s.router = newRouter()
	if s.config.Service.Log {
		s.logger.Info("Router initialized")
	}
}
func (s *Service) InitServer() {
	s.server = newServer(s.router, "", s.config.Service.Server.HttpPort)
	if s.config.Service.Log {
		s.logger.Info("Server initialized", zap.Int("port", s.GetPort()))
	}
}
func (s *Service) Start() error {
	if err := s.server.Start(); err != nil {
		if s.config.Service.Log {
			s.logger.Info("Failed starting service", zap.Error(err))
		}
		return err
	}
	if s.config.Service.Log {
		s.logger.Info("Service started successfully")
	}
	return nil
}

func (s *Service) Stop() {
	s.server.Shutdown()
	if s.config.Service.Log {
		s.logger.Info("Stopped service successfully")
	}
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
