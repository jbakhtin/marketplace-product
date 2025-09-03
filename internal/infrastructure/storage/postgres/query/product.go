package query

const (
	GetBySKU = `
		SELECT * FROM PRODUCTS
		WHERE sku = $1
		LIMIT 1
	`

	GetSKUs = `
		SELECT sku FROM products
		WHERE sku >= $1
		LIMIT $2
	`
)
