package entity

type OrderDetail struct {
	ID       int
	OrderID  int
	GameID   int
	Quantity int
	Subtotal float64
}
