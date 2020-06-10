package service

import (
	"net/http"
)

type Endpoint struct {
	name    string
	handler http.HandlerFunc
	method  string
	path    string
}

func MakeEndpoint(name, method, path string, handler http.HandlerFunc) *Endpoint {
	return &Endpoint{
		name:    name,
		handler: handler,
		method:  method,
		path:    path,
	}
}
