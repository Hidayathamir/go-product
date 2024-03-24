package usecase

import "context"

// IStock contains abstraction of usecase stock.
type IStock interface {
	Update(ctx context.Context)
}
