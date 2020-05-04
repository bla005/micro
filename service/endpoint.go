package service

import "net/http"

type Endpoint func(http.ResponseWriter, *http.Request)
