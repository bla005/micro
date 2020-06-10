package service

import (
	"net/http"
)

type Endpoint struct {
	Name    string           `json:"name"`
	Handler http.HandlerFunc `json:"-"`
	Method  string           `json:"method"`
	Path    string           `json:"path"`
}

func NewEndpoint(name, method, path string, handler http.HandlerFunc) *Endpoint {
	return &Endpoint{
		Name:    name,
		Handler: handler,
		Method:  method,
		Path:    path,
	}
}
