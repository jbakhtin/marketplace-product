package mock

import (
	"context"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/stretchr/testify/mock"
)

type ProductRepository struct {
	mock.Mock
}

func (p ProductRepository) GetProductBySKU(ctx context.Context, SKU domain.SKU) (domain.Product, error) {
	args := p.Called(ctx, SKU)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (p ProductRepository) GetSKUList(ctx context.Context, startAfterSKU domain.SKU, count int) ([]domain.SKU, error) {
	args := p.Called(ctx, startAfterSKU, count)

	if args.Get(0) != nil {
		return args.Get(0).([]domain.SKU), args.Error(1)
	}
	return nil, args.Error(1)
}
