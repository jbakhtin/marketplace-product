package entities

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jbakhtin/marketplace-product/internal/modules/product/domain"
)

var _ sql.Scanner = &FingerPrint{}

type FingerPrint map[string]any

// Scan - метод, реализующий интерфейс sql.Scanner sql.Out
func (m *FingerPrint) Scan(value interface{}) error {
	// Проверяем, что значение является срезом байтов ([]uint8)
	if value == nil {
		*m = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, received %T", value)
	}

	// Декодируем срез байтов в map[string]interface{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	return nil
}

type Product struct {
	ID    *int         `json:"id,omitempty" db:"id"`
	SKU   domain.SKU   `json:"sku" db:"sku"`
	Name  domain.Name  `json:"name" db:"name"`
	Price domain.Price `json:"price" db:"price"`
}

//func NewEntity(model models.Session) Session {
//	return Product{
//		...
//	}
//}

func (s *Product) ToModel() domain.Product {
	return domain.Product{
		SKU:   s.SKU,
		Name:  s.Name,
		Price: s.Price,
	}
}
