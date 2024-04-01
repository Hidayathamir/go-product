package goproduct

import (
	"errors"
	"time"
)

// ReqProductSearch -.
type ReqProductSearch struct {
	Keyword string `json:"keyword"`
}

// Validate -.
func (r ReqProductSearch) Validate() error {
	if r.Keyword == "" {
		return errors.New("keyword can not be empty")
	}
	return nil
}

// ResProductSearch -.
type ResProductSearch struct {
	Products []ResProductDetail `json:"products"`
}

// ReqProductDetail -.
type ReqProductDetail struct {
	ID   int64  `json:"id"`
	SKU  string `json:"sku"`
	Slug string `json:"slug"`
}

// Validate -.
func (r ReqProductDetail) Validate() error {
	if r.ID == 0 && r.SKU == "" && r.Slug == "" {
		return errors.New("at least one of id, sku, or slug must be provided")
	}
	return nil
}

// ResProductDetail -.
type ResProductDetail struct {
	ID          int64     `json:"id"`
	SKU         string    `json:"sku"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int32     `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
