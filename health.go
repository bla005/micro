package service

import (
	"encoding/json"
	"net/http"
	"time"
)

type healthResponse struct {
	Durationn time.Duration
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	healthResp := &healthResponse{
		Durationn: time.Now().Sub(startTime),
	}
	// if duration > timeout in config....xx

	if err := json.NewEncoder(w).Encode(healthResp); err != nil {
		//caca
	}
}
