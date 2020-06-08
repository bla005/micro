package service

import "errors"

var (
	// ErrNilLogger
	// ErrNilLogger error = errors.New("logger is nil")
	//
	ErrNilConfig error = errors.New("config is nil")
	ErrNilRouter error = errors.New("router is nil")
)
