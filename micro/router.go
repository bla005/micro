package micro

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func newRouter(endpoints []*Endpoint) http.Handler {

	router := httprouter.New()
	/*var (
		userRepository = postgres
		userStore      = redis
		userService    = services.NewUserService(userRepository, userStore, auth)
		userEndpoints  = endpoints.NewUserEndpoints(userService, logger, framework)
	)*/

	/*
		router.HandlerFunc("POST", "/signup", userEndpoints.Signup)
		router.HandlerFunc("POST", "/login", userEndpoints.Login)
		router.HandlerFunc("POST", "/confirm", userEndpoints.ConfirmAccount)
		router.HandlerFunc("GET", "/user", userEndpoints.User)
	*/

	router.HandlerFunc(endpoints[0].Method, endpoints[0].Path, endpoints[0].Handler)

	return router
}
