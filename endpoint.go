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

// Name
func (e *Endpoint) Name() string {
	return e.name
}

// Handler
func (e *Endpoint) Handler() http.HandlerFunc {
	return e.handler
}

// Method
func (e *Endpoint) Method() string {
	return e.method
}

// Path
func (e *Endpoint) Path() string {
	return e.path
}

// MakeEndpoint
func MakeEndpoint(name, method, path string, handler http.HandlerFunc) *Endpoint {
	return &Endpoint{
		name:    name,
		handler: handler,
		method:  method,
		path:    path,
	}
}
