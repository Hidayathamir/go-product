package http

import (
	"github.com/Hidayathamir/go-product/pkg/goproductdto"
)

// ResSearch -.
type ResSearch struct {
	Data  goproductdto.ResProductSearch `json:"data"`
	Error any                           `json:"error"`
}

// ResGetDetail -.
type ResGetDetail struct {
	Data  goproductdto.ResProductDetail `json:"data"`
	Error any                           `json:"error"`
}
