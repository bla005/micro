package micro

import "net/http"

type Endpoint struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}
