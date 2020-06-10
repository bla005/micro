package service

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	uptime       string            `json:"uptime"`
	dependencies map[string]string `json:"dependencies"`
}

func newHealthResponse(uptime string, dependencies map[string]string) *healthResponse {
	return &healthResponse{uptime: uptime, dependencies: dependencies}
}

func (s *Service) healthHandler(w http.ResponseWriter, r *http.Request) {
	health := s.Health()
	resp := newHealthResponse(s.Uptime().String(), health)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
}
