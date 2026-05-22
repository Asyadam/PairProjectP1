package repository

import (
	"database/sql"
	"errors"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(userID int, gameID int, quantity int) error {
	tx, err := r.DB.Begin()

	if err != nil {
		return err
	}

	var price float64
	var stock int

	err = tx.QueryRow(`
		SELECT price, stock
		FROM games
		WHERE id = ?
	`, gameID).Scan(&price, &stock)

	if err != nil {
		tx.Rollback()
		return err
	}

	if stock < quantity {
		tx.Rollback()
		return errors.New("stock is not enough")
	}

	subtotal := price * float64(quantity)

	result, err := tx.Exec(`
		INSERT INTO orders(user_id, total_price, status)
		VALUES (?, ?, ?)
	`, userID, subtotal, "paid")

	if err != nil {
		tx.Rollback()
		return err
	}

	orderID, err := result.LastInsertId()

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO order_details(order_id, game_id, quantity, subtotal)
		VALUES (?, ?, ?, ?)
	`, orderID, gameID, quantity, subtotal)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		UPDATE games
		SET stock = stock - ?
		WHERE id = ?
	`, quantity, gameID)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
