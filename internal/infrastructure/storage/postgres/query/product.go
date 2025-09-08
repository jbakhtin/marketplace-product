package query

const (
	GetBySKU = `
		SELECT id, sku, name, price, created_at, updated_at FROM PRODUCTS
		WHERE sku = $1
		LIMIT 1
	`

	GetSKUs = `
		SELECT sku FROM products
		WHERE sku > $1
		LIMIT $2
	`
)
