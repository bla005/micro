package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router interface {
	HandlerFunc(string, string, http.HandlerFunc)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func newRouter() Router {
	return httprouter.New()
}
