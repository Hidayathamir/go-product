package ctrlgrpc

import (
	"context"
	"testing"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/usecase/interfaces/mock"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUnitStockIncrementStock(t *testing.T) {
	t.Parallel()

	t.Run("increment success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseStock := mock.NewMockUsecaseStock(ctrl)

		s := &Stock{
			cfg:          config.Config{},
			usecaseStock: usecaseStock,
		}

		reqUsecase := goproduct.ReqIncrementStock{ProductID: 123123}

		usecaseStock.EXPECT().IncrementStock(context.Background(), reqUsecase).Return(nil)

		req := &goproductgrpc.ReqIncrementStock{ProductId: reqUsecase.ProductID}

		res, err := s.IncrementStock(context.Background(), req)

		require.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("usecase error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseStock := mock.NewMockUsecaseStock(ctrl)

		s := &Stock{
			cfg:          config.Config{},
			usecaseStock: usecaseStock,
		}

		reqUsecase := goproduct.ReqIncrementStock{ProductID: 123123}

		usecaseStock.EXPECT().IncrementStock(context.Background(), reqUsecase).Return(assert.AnError)

		req := &goproductgrpc.ReqIncrementStock{ProductId: reqUsecase.ProductID}

		res, err := s.IncrementStock(context.Background(), req)

		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}

func TestUnitStockDecrementStock(t *testing.T) {
	t.Parallel()

	t.Run("decrement success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseStock := mock.NewMockUsecaseStock(ctrl)

		s := &Stock{
			cfg:          config.Config{},
			usecaseStock: usecaseStock,
		}

		reqUsecase := goproduct.ReqDecrementStock{ProductID: 23423}

		usecaseStock.EXPECT().DecrementStock(context.Background(), reqUsecase).Return(nil)

		req := &goproductgrpc.ReqDecrementStock{ProductId: reqUsecase.ProductID}

		res, err := s.DecrementStock(context.Background(), req)

		require.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("usecase error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseStock := mock.NewMockUsecaseStock(ctrl)

		s := &Stock{
			cfg:          config.Config{},
			usecaseStock: usecaseStock,
		}

		reqUsecase := goproduct.ReqDecrementStock{ProductID: 23423}

		usecaseStock.EXPECT().DecrementStock(context.Background(), reqUsecase).Return(assert.AnError)

		req := &goproductgrpc.ReqDecrementStock{ProductId: reqUsecase.ProductID}

		res, err := s.DecrementStock(context.Background(), req)

		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
