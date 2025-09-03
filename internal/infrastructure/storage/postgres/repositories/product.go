package repositories

import (
	"context"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
)

type CartStorage struct {
}

func NewProductStorage() (CartStorage, error) {
	return CartStorage{}, nil
}

func (c *CartStorage) GetProductBySKU(ctx context.Context, SKU domain.SKU) (domain.Product, error) {
	return domain.Product{}, nil
}

func (c *CartStorage) GetSKUList(ctx context.Context, startSKU domain.SKU, count int) ([]domain.SKU, error) {
	return []domain.SKU{}, nil
}
