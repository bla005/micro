package service

import (
	"encoding/json"
	"net/http"
	"time"
)

func (s *service) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Uptime time.Duration `json:"uptime"`
		Check  string        `json:"check"`
	}{
		Uptime: s.GetUptime(),
		Check:  "OK",
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
