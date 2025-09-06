package product

import (
	"github.com/jbakhtin/marketplace-product/internal/modules/product/ports"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/use_case"
)

type Module struct {
	productUseCase use_case.ProductUseCase
}

func InitModule(logger ports.Logger, productRepository ports.ProductRepository) (Module, error) {
	productUseCase, err := use_case.NewProductUseCase(logger, productRepository)
	if err != nil {
		return Module{}, err
	}

	return Module{
		productUseCase: productUseCase,
	}, nil
}

func (m Module) GetProductUseCase() use_case.ProductUseCase {
	return m.productUseCase
}
