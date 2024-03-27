package usecase

import (
	"context"
	"testing"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/usecase/interfaces/mock"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUnitStockIncrementStock(t *testing.T) {
	t.Parallel()

	t.Run("increment stock success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoStock := mock.NewMockRepoStock(ctrl)

		s := &Stock{
			cfg:       config.Config{},
			repoStock: repoStock,
		}

		req := goproduct.ReqIncrementStock{
			ProductID: 23432,
		}

		repoStock.EXPECT().
			IncrementStock(context.Background(), req.ProductID).
			Return(nil)

		err := s.IncrementStock(context.Background(), req)

		require.NoError(t, err)
	})
	t.Run("increment stock error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoStock := mock.NewMockRepoStock(ctrl)

		s := &Stock{
			cfg:       config.Config{},
			repoStock: repoStock,
		}

		req := goproduct.ReqIncrementStock{
			ProductID: 23432,
		}

		repoStock.EXPECT().
			IncrementStock(context.Background(), req.ProductID).
			Return(assert.AnError)

		err := s.IncrementStock(context.Background(), req)

		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("request validate error should return error", func(t *testing.T) {
		t.Parallel()

		t.Run("product id empty should return error", func(t *testing.T) {
			t.Parallel()

			s := &Stock{
				cfg: config.Config{},
			}

			req := goproduct.ReqIncrementStock{}

			err := s.IncrementStock(context.Background(), req)

			require.Error(t, err)
			require.ErrorIs(t, err, goproduct.ErrRequestInvalid)
		})
	})
}

func TestUnitStockDecrementStock(t *testing.T) {
	t.Parallel()

	t.Run("decrement stock success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoStock := mock.NewMockRepoStock(ctrl)

		s := &Stock{
			cfg:       config.Config{},
			repoStock: repoStock,
		}

		req := goproduct.ReqDecrementStock{
			ProductID: 2342,
		}

		repoStock.EXPECT().
			DecrementStock(context.Background(), req.ProductID).
			Return(nil)

		err := s.DecrementStock(context.Background(), req)

		require.NoError(t, err)
	})
	t.Run("decrement stock error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoStock := mock.NewMockRepoStock(ctrl)

		s := &Stock{
			cfg:       config.Config{},
			repoStock: repoStock,
		}

		req := goproduct.ReqDecrementStock{
			ProductID: 2342,
		}

		repoStock.EXPECT().
			DecrementStock(context.Background(), req.ProductID).
			Return(assert.AnError)

		err := s.DecrementStock(context.Background(), req)

		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("request validate error should return error", func(t *testing.T) {
		t.Parallel()

		t.Run("product id empty should return error", func(t *testing.T) {
			t.Parallel()

			s := &Stock{
				cfg: config.Config{},
			}

			req := goproduct.ReqDecrementStock{}

			err := s.DecrementStock(context.Background(), req)

			require.Error(t, err)
			require.ErrorIs(t, err, goproduct.ErrRequestInvalid)
		})
	})
}
