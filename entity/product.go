package entity

type Product struct {
	ID     int
	Name   string
	Price  string
	Stocks int
}

func (p Product) StockStatus() string {
	if p.Stocks > 0 {
		return "In Stock"
	}
	return "Out of Stock"
}
