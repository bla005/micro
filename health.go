package service

import (
	"encoding/json"
	"net/http"
)

func (s *service) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Uptime string `json:"uptime"`
		Check  string `json:"check"`
	}{
		Uptime: s.GetUptime().String(),
		Check:  "ok",
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
