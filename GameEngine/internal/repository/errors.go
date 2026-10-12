package repository

import "errors"

var (
	ErrNotFound       = errors.New("record not found")
	ErrInvalidInput   = errors.New("invalid repository input")
	ErrInvalidJSON    = errors.New("invalid JSON")
)
