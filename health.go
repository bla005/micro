package service

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

type healthResponse struct {
	Durationn time.Duration
	Pula      int
}

func checkHealth() int {
	conn, err := net.DialTimeout("tcp", "localhost:8088/health", 2*time.Second)
	if err != nil {
		return 0
	}
	defer conn.Close()
	return 1
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	healthResp := &healthResponse{
		Durationn: time.Now().Sub(startTime),
		Pula:      checkHealth(),
	}

	// if duration > timeout in config....xx

	if err := json.NewEncoder(w).Encode(healthResp); err != nil {
		//caca
	}
}
