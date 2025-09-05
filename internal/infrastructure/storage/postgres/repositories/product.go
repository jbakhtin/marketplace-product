package repositories

import (
	"context"
	"database/sql"

	"github.com/jbakhtin/marketplace-product/internal/infrastructure/storage/postgres/entities"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/storage/postgres/query"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/ports"
	"github.com/pkg/errors"
)

type ProductStorage struct {
	db *sql.DB
}

func NewProductStorage(db *sql.DB) ports.ProductRepository {
	return &ProductStorage{
		db: db,
	}
}

func (p *ProductStorage) GetProductBySKU(ctx context.Context, SKU domain.SKU) (domain.Product, error) {
	var product entities.Product

	err := p.db.QueryRowContext(ctx, query.GetBySKU, SKU).Scan(
		&product.ID,
		&product.SKU,
		&product.Name,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Product{}, errors.New("product not found")
		}
		return domain.Product{}, errors.Wrap(err, "failed to get product by SKU")
	}

	return product.ToModel(), nil
}

func (p *ProductStorage) GetSKUList(ctx context.Context, startAfterSKU domain.SKU, count int) ([]domain.SKU, error) {
	rows, err := p.db.QueryContext(ctx, query.GetSKUs, startAfterSKU, count)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get SKU list")
	}
	defer rows.Close()

	var skus []domain.SKU
	for rows.Next() {
		var sku domain.SKU
		if err := rows.Scan(&sku); err != nil {
			return nil, errors.Wrap(err, "failed to scan SKU")
		}
		skus = append(skus, sku)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error during rows iteration")
	}

	return skus, nil
}
