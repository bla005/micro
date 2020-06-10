package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var DefaultTLSConfig = &tls.Config{
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

func start(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("service: failed listening")
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)

	if err := srv.Shutdown(ctx); err != nil {
		panic("shutdown failed")
	}
}

func shutdown(srv *http.Server) {
	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(context.Background()); err != nil {
		panic("shutdown failed")
	}
}

func NewServer(handler http.Handler, config *Config) *http.Server {
	if config == nil {
		panic("nil config")
	}
	addr := fmt.Sprintf("%s:%d", config.Service.Server.Host, config.Service.Server.Port)
	srv := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       config.Service.Server.Timeout.Read * time.Second,
		WriteTimeout:      config.Service.Server.Timeout.Write * time.Second,
		IdleTimeout:       config.Service.Server.Timeout.Idle * time.Second,
		ReadHeaderTimeout: config.Service.Server.Timeout.ReadHeader * time.Second,
		TLSConfig:         DefaultTLSConfig,
	}
	return srv
}
