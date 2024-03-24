package usecase

import "context"

// IProduct contains abstraction of usecase product.
type IProduct interface {
	GetAll(ctx context.Context)
	GetDetail(ctx context.Context)
}
