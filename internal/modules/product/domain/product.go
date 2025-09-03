package domain

type SKU int32
type Name string
type Price int32

type Product struct {
	Name  Name
	SKU   SKU
	Price Price
}
