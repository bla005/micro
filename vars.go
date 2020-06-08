package service

import (
	"crypto/tls"
	"time"
)

var (
	// startTime is used to calculate the service uptime
	startTime *time.Time

	// defaultTLSConfig is the default TLS config used by the server
	defaultTLSConfig = &tls.Config{
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

	defaultServiceConfig = &Config{
		Service: struct {
			LogDir string "yaml:\"logdir\""
			Health struct {
				Path string "yaml:\"path\""
			} "yaml:\"health\""
			Server struct {
				Host    string "yaml:\"host\""
				Port    int    "yaml:\"port\""
				Ssl     bool   "yaml:\"ssl\""
				Timeout struct {
					Read       time.Duration "yaml:\"read\""
					Write      time.Duration "yaml:\"write\""
					Idle       time.Duration "yaml:\"idle\""
					ReadHeader time.Duration "yaml:\"read_header\""
				} "yaml:\"timeout\""
			} "yaml:\"server\""
		}{
			LogDir: "logs",
			Health: struct {
				Path string "yaml:\"path\""
			}{
				Path: "/health",
			},
			Server: struct {
				Host    string "yaml:\"host\""
				Port    int    "yaml:\"port\""
				Ssl     bool   "yaml:\"ssl\""
				Timeout struct {
					Read       time.Duration "yaml:\"read\""
					Write      time.Duration "yaml:\"write\""
					Idle       time.Duration "yaml:\"idle\""
					ReadHeader time.Duration "yaml:\"read_header\""
				} "yaml:\"timeout\""
			}{
				Host: "",
				Port: 8088,
				Ssl:  false,
				Timeout: struct {
					Read       time.Duration "yaml:\"read\""
					Write      time.Duration "yaml:\"write\""
					Idle       time.Duration "yaml:\"idle\""
					ReadHeader time.Duration "yaml:\"read_header\""
				}{
					Idle:       5 * time.Second,
					Read:       5 * time.Second,
					Write:      5 * time.Second,
					ReadHeader: 5 * time.Second,
				},
			},
		},
	}
)
