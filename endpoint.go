package service

import "net/http"

type Endpoint struct {
	Handler http.HandlerFunc
	Method  string
	Path    string
}

func newEndpoint(handler http.HandlerFunc, method, path string) *Endpoint {
	return &Endpoint{
		Handler: handler,
		Method:  method,
		Path:    path,
	}
}
