package repository

import (
	"database/sql"
	"fmt"
)

type ReportRepository struct {
	DB *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (r *ReportRepository) UserReport() error {
	rows, err := r.DB.Query(`
		SELECT id, email, role, created_at
		FROM users
	`)

	if err != nil {
		return err
	}

	defer rows.Close()

	fmt.Println("==== USER REPORT ====")

	for rows.Next() {
		var id int
		var email string
		var role string
		var createdAt string

		err := rows.Scan(&id, &email, &role, &createdAt)

		if err != nil {
			return err
		}

		fmt.Println("ID:", id)
		fmt.Println("Email:", email)
		fmt.Println("Role:", role)
		fmt.Println("Created At:", createdAt)
		fmt.Println("=====================")
	}

	return nil
}

func (r *ReportRepository) SalesReport() error {
	rows, err := r.DB.Query(`
		SELECT 
			orders.id,
			users.email,
			orders.total_price,
			orders.status,
			orders.order_date
		FROM orders
		JOIN users ON orders.user_id = users.id
	`)

	if err != nil {
		return err
	}

	defer rows.Close()

	fmt.Println("==== SALES REPORT ====")

	for rows.Next() {
		var orderID int
		var email string
		var totalPrice float64
		var status string
		var orderDate string

		err := rows.Scan(
			&orderID,
			&email,
			&totalPrice,
			&status,
			&orderDate,
		)

		if err != nil {
			return err
		}

		fmt.Println("Order ID:", orderID)
		fmt.Println("User Email:", email)
		fmt.Println("Total Price:", totalPrice)
		fmt.Println("Status:", status)
		fmt.Println("Order Date:", orderDate)
		fmt.Println("======================")
	}

	return nil
}

func (r *ReportRepository) StockReport() error {
	rows, err := r.DB.Query(`
		SELECT id, title, stock
		FROM games
	`)

	if err != nil {
		return err
	}

	defer rows.Close()

	fmt.Println("==== STOCK REPORT ====")

	for rows.Next() {
		var id int
		var title string
		var stock int

		err := rows.Scan(&id, &title, &stock)

		if err != nil {
			return err
		}

		fmt.Println("Game ID:", id)
		fmt.Println("Title:", title)
		fmt.Println("Stock:", stock)
		fmt.Println("======================")
	}

	return nil
}
