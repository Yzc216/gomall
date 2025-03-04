package types

type OrderItem struct {
	ProductName string
	Picture     string
	Qty         uint32
	Cost        float64
}

type Order struct {
	OrderId     string
	CreatedDate string
	Cost        float64
	Items       []OrderItem
}
