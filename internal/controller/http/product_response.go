package http

import "github.com/Hidayathamir/go-product/pkg/goproduct"

// ResSearch -.
type ResSearch struct {
	Data  goproduct.ResProductSearch `json:"data"`
	Error any                        `json:"error"`
}

// ResGetDetail -.
type ResGetDetail struct {
	Data  goproduct.ResProductDetail `json:"data"`
	Error any                        `json:"error"`
}
