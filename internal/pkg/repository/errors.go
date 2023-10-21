package repository

import "errors"

var (
	// ErrObjectNotFound is used when nothing was retrieved from database
	ErrObjectNotFound = errors.New("object not found")
)
