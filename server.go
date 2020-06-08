package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	server *http.Server
	cert   string
	key    string
	host   string
	port   int
	// done   chan (os.Signal)
}

func (s *server) setTLSConfig(config *tls.Config) {
	defaultTLSConfig = config
}
func (s *server) getTLSConfig() *tls.Config {
	return defaultTLSConfig
}
func (s *server) getHost() string {
	return s.host
}
func (s *server) getPort() int {
	return s.port
}
func (s *server) start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Fatalf("listenAndServe: %v", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// wait for channel to receive
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// s.server.SetKeepAlivesEnabled(false)
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown: %v", err)
	}

	close(done)
	// return nil
}
func (s *server) startWithTLS() error {
	if err := s.server.ListenAndServeTLS(s.cert, s.key); err != nil {
		return err
	}
	return nil
}
func (s *server) shutdown() {
	// done := make(chan os.Signal, 1)
	// signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// <-done

	// signal.Notify(s.done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// <-s.done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// s.server.SetKeepAlivesEnabled(false)
	if err := s.server.Shutdown(ctx); err != nil {
		// log.Fatalf("Failed shutting down server: %v", err)
		os.Exit(0)
	}
	// close(done)
	// close(s.done)
}

func newServer(handler http.Handler, config *Config) *server {
	addr := fmt.Sprintf("%s:%d", config.Service.Server.Host, config.Service.Server.Port)
	s := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       config.Service.Server.Timeout.Read * time.Second,
		WriteTimeout:      config.Service.Server.Timeout.Write * time.Second,
		IdleTimeout:       config.Service.Server.Timeout.Idle * time.Second,
		ReadHeaderTimeout: config.Service.Server.Timeout.ReadHeader * time.Second,
		TLSConfig:         defaultTLSConfig,
	}
	return &server{
		server: s,
		host:   config.Service.Server.Host,
		port:   config.Service.Server.Port,
		// done:   make(chan os.Signal, 1),
	}
}
