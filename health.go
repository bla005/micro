package service

import (
	"encoding/json"
	"net/http"
)

func (s *service) Health() map[string]string {
	e := make(map[string]string)
	for i := 0; i < len(s.dependencies); i++ {
		if err := s.dependencies[i].Ping(); err != nil {
			e[s.dependencies[i].Name] = "critical"
		}
	}
	return e
}

func (s *service) healthHandler(w http.ResponseWriter, r *http.Request) {
	e := s.Health()
	resp := struct {
		Uptime    string            `json:"uptime"`
		Externals map[string]string `json:"externals"`
	}{
		Uptime:    s.Uptime().String(),
		Externals: e,
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
