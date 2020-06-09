package service

import "errors"

var (
	ErrNilConfig error = errors.New("config is nil")
	ErrNilServer error = errors.New("server is nil")
	ErrNilRouter error = errors.New("router is nil")
)
