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

type Server interface {
	Start() error
	StartWithTLS() error
	GetHost() string
	GetPort() int
	Shutdown()
}

type server struct {
	server *http.Server
	cert   string
	key    string
	host   string
	port   int
}

func (s *server) Start() error {
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
func (s *server) GetHost() string {
	return s.host
}
func (s *server) GetPort() int {
	return s.port
}
func (s *server) StartWithTLS() error {
	if err := s.server.ListenAndServeTLS(s.cert, s.key); err != nil {
		return err
	}
	return nil
}
func (s *server) Shutdown() {
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

func newServer(handler http.Handler, host string, port int) Server {
	addr := fmt.Sprintf("%s:%d", host, port)
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
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		TLSConfig:         tlsConfig,
	}
	return &server{server: s, host: host, port: port}
}
