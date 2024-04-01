package goproduct

import "errors"

// ReqIncrementStock -.
type ReqIncrementStock struct {
	ProductID int64 `json:"product_id"`
}

// Validate -.
func (r ReqIncrementStock) Validate() error {
	if r.ProductID == 0 {
		return errors.New("product_id can not be empty")
	}
	return nil
}

// ReqDecrementStock -.
type ReqDecrementStock struct {
	ProductID int64 `json:"product_id"`
}

// Validate -.
func (r ReqDecrementStock) Validate() error {
	if r.ProductID == 0 {
		return errors.New("product_id can not be empty")
	}
	return nil
}
