package service

import "errors"

var (
	ErrNilConfig      error = errors.New("config is nil")
	ErrNotInitialized error = errors.New("service not initialized")
)
