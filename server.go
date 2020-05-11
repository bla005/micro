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
}

func (s *server) getHost() string {
	return s.host
}
func (s *server) getPort() int {
	return s.port
}
func (s *server) start() error {
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
func (s *server) startWithTLS() error {
	if err := s.server.ListenAndServeTLS(s.cert, s.key); err != nil {
		return err
	}
	return nil
}
func (s *server) shutdown() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.server.SetKeepAlivesEnabled(false)
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed shutting down server: %v", err)
	}
	close(done)
}

func newServer(handler http.Handler, config *Config) *server {
	addr := fmt.Sprintf("%s:%d", config.Service.Server.Host, config.Service.Server.Port)
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{
			tls.CurveP521,
			tls.CurveP384,
			tls.CurveP256,
		},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	s := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       config.Service.Server.Timeout.Read * time.Second,
		WriteTimeout:      config.Service.Server.Timeout.Write * time.Second,
		IdleTimeout:       config.Service.Server.Timeout.Idle * time.Second,
		ReadHeaderTimeout: config.Service.Server.Timeout.ReadHeader * time.Second,
		TLSConfig:         tlsConfig,
	}
	return &server{server: s, host: config.Service.Server.Host, port: config.Service.Server.Port}
}
