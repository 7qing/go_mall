package types

type OrderItem struct {
	ProductName string
	Picture     string
	Qty         int32
	Cost        float32
}

type Order struct {
	OrderId    string
	CreateDate string
	Cost       float32
	Items      []OrderItem
}
