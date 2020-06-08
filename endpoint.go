package service

import "net/http"

// Endpoint object
type Endpoint struct {
	Name    string
	Handler http.HandlerFunc
	Method  string
	Path    string
}

func newEndpoint(name string, handler http.HandlerFunc, method, path string) *Endpoint {
	return &Endpoint{
		Name:    name,
		Handler: handler,
		Method:  method,
		Path:    path,
	}
}

// Register an endpoint
func (s *service) RegisterEndpoint(name string, handler http.HandlerFunc, method, path string) {
	endpoint := newEndpoint(name, handler, method, path)
	s.endpoints = append(s.endpoints, endpoint)
}

// Register all endpoints
func (s *service) RegisterEndpoints() {
	for i := 0; i < len(s.endpoints); i++ {
		s.router.HandlerFunc(s.endpoints[i].Method, s.endpoints[i].Path, s.endpoints[i].Handler)
	}
}
