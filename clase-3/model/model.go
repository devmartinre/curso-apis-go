package model

import "errors"

var (
	// ErrPersonCanNotBeNil .
	ErrPersonCanNotBeNil = errors.New("person cannot be nil")

	// ErrIDPersonDoesNotExists .
	ErrIDPersonDoesNotExists = errors.New("person does not exists")
)