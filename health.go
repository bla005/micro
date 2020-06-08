package service

import (
	"encoding/json"
	"net/http"
)

func (s *service) healthHandler(w http.ResponseWriter, r *http.Request) {
	// check store
	// check database
	resp := struct {
		Uptime        string `json:"uptime"`
		Db            string `json:"database"`
		Store         string `json:"store"`
		GeneralHealth string `json:"general_health"`
	}{
		Uptime:        s.GetUptime().String(),
		Db:            "ok",
		Store:         "ok",
		GeneralHealth: "ok",
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
