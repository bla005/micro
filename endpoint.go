package service

import (
	"net/http"
)

type Endpoint struct {
	Name    string `json:"name"`
	handler http.HandlerFunc
	Method  string `json:"method"`
	Path    string `json:"path"`
}

func MakeEndpoint(name, method, path string, handler http.HandlerFunc) *Endpoint {
	return &Endpoint{
		Name:    name,
		handler: handler,
		Method:  method,
		Path:    path,
	}
}
