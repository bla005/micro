package service

import "net/http"

type Endpoint struct {
	Handler http.HandlerFunc
	Method  string
	Path    string
}
