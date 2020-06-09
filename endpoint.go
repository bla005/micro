package service

import (
	"net/http"
)

// Endpoint structure
type Endpoint struct {
	name    string
	handler http.HandlerFunc
	method  string
	path    string
}

func (e *Endpoint) Name() string {
	return e.name
}
func (e *Endpoint) Handler() http.HandlerFunc {
	return e.handler
}
func (e *Endpoint) Method() string {
	return e.method
}
func (e *Endpoint) Path() string {
	return e.path
}

func MakeEndpoint(name method, path string, handler http.HandlerFunc) *Endpoint {
	return &Endpoint{
		name:    name,
		handler: handler,
		method:  method,
		path:    path,
	}
}
