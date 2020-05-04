package service

import "github.com/julienschmidt/httprouter"

func newRouter() *httprouter.Router {
	return httprouter.New()
}
