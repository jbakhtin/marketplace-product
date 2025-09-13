package product

import (
	"github.com/jbakhtin/marketplace-product/internal/modules/product/ports"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/use_case"
)

type Config interface {
}

type Handler struct {
	cfg     Config
	log     ports.Logger
	useCase use_case.UseCaseInterface
}

func NewProductHandler(cfg Config, logger ports.Logger, useCase use_case.UseCaseInterface) (Handler, error) {
	return Handler{
		cfg:     cfg,
		log:     logger,
		useCase: useCase,
	}, nil
}
