package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// type Server interface {
// 	start()
// 	startWithTLS()
// 	Shutdown()
//
// 	Host() string
// 	Port() int
//
// 	LoadCert(file string) error
// 	LoadKey(file string) error
//
// 	TLSConfig() *tls.Config
// 	SetTLSConfig(config *tls.Config)
// }

type Server struct {
	srv  *http.Server
	cert string
	key  string
	host string
	port int
}

func (s *Server) LoadCert(file string) error {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	s.cert = string(contents)
	return nil
}
func (s *Server) LoadKey(file string) error {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	s.key = string(contents)
	return nil
}

func (s *Server) SetTLSConfig(config *tls.Config) {
	defaultTLSConfig = config
}
func (s *Server) TLSConfig() *tls.Config {
	return defaultTLSConfig
}
func (s *Server) Host() string {
	return s.host
}
func (s *Server) Port() int {
	return s.port
}
func (s *Server) start() {
	// Server must run in a different goroutine so it doesn't block
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// log.Fatalf("listenAndServe: %v", err)
			panic("failed listenAndServe")
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Wait for done to receive a signal
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Keep-alives can be disabled
	s.srv.SetKeepAlivesEnabled(false)

	if err := s.srv.Shutdown(ctx); err != nil {
		// log.Fatalf("shutdown: %v", err)
		panic("failed shutting down server")
	}
	close(done)
}
func (s *Server) startWithTLS() {
	if err := s.srv.ListenAndServeTLS(s.cert, s.key); err != nil {
		panic("failed starting server with tls")
	}
}
func (s *Server) shutdown() {
	s.srv.SetKeepAlivesEnabled(false)
	if err := s.srv.Shutdown(context.Background()); err != nil {
		panic("failed shutting down")
	}
}

func NewServer(handler http.Handler, config *Config) *Server {
	addr := fmt.Sprintf("%s:%d", config.ServerHost(), config.ServerPort())
	s := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       config.ServerReadTimeout() * time.Second,
		WriteTimeout:      config.ServerWriteTimeout() * time.Second,
		IdleTimeout:       config.ServerIdleTimeout() * time.Second,
		ReadHeaderTimeout: config.ServerReadHeaderTimeout() * time.Second,
		TLSConfig:         defaultTLSConfig,
	}
	return &Server{
		srv:  s,
		host: config.ServerHost(),
		port: config.ServerPort(),
		cert: "",
		key:  "",
	}
}
