package types

type Consignee struct {
	Email string

	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	Consignee   Consignee
	OrderId     uint64
	CreatedDate string
	OrderState  string
	Cost        float64
	Items       []OrderItem
}

type OrderItem struct {
	SpuId       uint64
	SkuId       uint64
	ProductName string
	Picture     string
	Qty         uint32
	Cost        float64
}
