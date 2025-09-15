package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/ports"
)

type UseCaseInterface interface {
	GetProductBySKU(ctx context.Context, SKU domain.SKU) (domain.Product, error)
	GetSKUList(ctx context.Context, startAfterSKU domain.SKU, count int) ([]domain.SKU, error)
}

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

func (o *ProductUseCase) GetProductBySKU(ctx context.Context, SKU domain.SKU) (domain.Product, error) {
	product, err := o.productRepository.GetProductBySKU(ctx, SKU)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (o *ProductUseCase) GetSKUList(ctx context.Context, startAfterSKU domain.SKU, count int) ([]domain.SKU, error) {
	skus, err := o.productRepository.GetSKUList(ctx, startAfterSKU, count)
	if err != nil {
		return nil, err
	}

	return skus, nil
}
