package product

import (
	"context"

	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetProductBySKU(ctx context.Context, SKU domain.SKU) (domain.Product, error) {
	args := m.Called(ctx, SKU)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductService) GetSKUList(ctx context.Context, startSKU domain.SKU, count int) ([]domain.SKU, error) {
	args := m.Called(ctx, startSKU, count)

	if args.Get(0) != nil {
		return args.Get(0).([]domain.SKU), args.Error(1)
	}
	return nil, args.Error(1)
}
