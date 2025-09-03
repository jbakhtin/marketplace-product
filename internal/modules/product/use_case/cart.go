package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/ports"
)

type ProductUseCase struct {
	logger            ports.Logger
	productRepository ports.ProductRepository
}

func NewProductUseCase(
	logger ports.Logger,
	productRepository ports.ProductRepository,
) (ProductUseCase, error) {
	return ProductUseCase{
		logger:            logger,
		productRepository: productRepository,
	}, nil
}

func (o *ProductUseCase) GetProductUseCase(ctx context.Context, SKU domain.SKU) (domain.Product, error) {
	return domain.Product{}, nil
}
