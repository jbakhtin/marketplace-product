package product

import (
	openapi "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/openapi/v1"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
)

func toOpenAPIProduct(product domain.Product) openapi.Product {
	name := string(product.Name)
	sku := int32(product.SKU)
	price := int32(product.Price)

	return openapi.Product{
		Name:  &name,
		SKU:   &sku,
		Price: &price,
	}
}

func toOpenAPISKUs(skus []domain.SKU) []int32 {
	result := make([]int32, len(skus))
	for i, sku := range skus {
		result[i] = int32(sku)
	}

	return result
}
