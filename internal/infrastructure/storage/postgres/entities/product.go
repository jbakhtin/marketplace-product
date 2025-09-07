package entities

import (
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
)

type Product struct {
	ID        *int             `json:"id,omitempty" db:"id"`
	SKU       domain.SKU       `json:"sku" db:"sku"`
	Name      domain.Name      `json:"name" db:"name"`
	Price     domain.Price     `json:"price" db:"price"`
	CreatedAt pgtype.Timestamp `json:"created_at" db:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at" db:"updated_at"`
}

func (s *Product) ToModel() domain.Product {
	return domain.Product{
		SKU:   s.SKU,
		Name:  s.Name,
		Price: s.Price,
	}
}
