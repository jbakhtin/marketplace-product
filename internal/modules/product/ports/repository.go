package ports

import (
	"context"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
)

type ProductRepository interface {
	GetProductBySKU(ctx context.Context, SKU domain.SKU) (domain.Product, error)
	GetSKUList(ctx context.Context, startAfterSKU domain.SKU, count int) ([]domain.SKU, error)
}
