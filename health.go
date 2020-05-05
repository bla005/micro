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
	timeStart := time.Now()
	healthResp := &healthResponse{
		Durationn: time.Since(timeStart),
	}
	// if duration > timeout in config....

	if err := json.NewEncoder(w).Encode(healthResp); err != nil {
		//caca
	}
}
