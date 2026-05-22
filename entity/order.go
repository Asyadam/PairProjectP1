package entity

type Order struct {
	ID         int
	UserID     int
	OrderDate  string
	TotalPrice float64
	Status     string
}
