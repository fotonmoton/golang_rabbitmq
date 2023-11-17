package warehouse

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Qty    int    `json:"qty"`
	Active bool   `json:"active"`
}

func NewProduct(name string, qty int) Product {
	return Product{
		ID:     0,
		Name:   name,
		Qty:    qty,
		Active: true,
	}
}
