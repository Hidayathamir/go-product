package goproducterror

import "errors"

var (
	// ErrProductNotFound occurs when product does not exists.
	ErrProductNotFound = errors.New("product not found")
	// ErrProductDuplicateSKU occurs when product sku duplicate.
	ErrProductDuplicateSKU = errors.New("product sku duplicate")
	// ErrProductDuplicateSlug occurs when product slug duplicate.
	ErrProductDuplicateSlug = errors.New("product slug duplicate")
)
