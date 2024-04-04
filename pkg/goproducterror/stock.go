package goproducterror

import "errors"

// ErrStockDuplicateProductID occurs when stock product id duplciate.
var ErrStockDuplicateProductID = errors.New("stock duplicate product id")
