package goproduct

import "errors"

var (
	// ErrRequestInvalid occurs when request invalid.
	ErrRequestInvalid = errors.New("request invalid")
	// ErrProductNotFound occurs when product does not exists.
	ErrProductNotFound = errors.New("product not found")
)
