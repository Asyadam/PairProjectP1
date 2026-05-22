package handler

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/repository"
)

type OrderHandler struct {
	OrderRepo *repository.OrderRepository
}

func NewOrderHandler(orderRepo *repository.OrderRepository) *OrderHandler {
	return &OrderHandler{OrderRepo: orderRepo}
}

func (h *OrderHandler) CreateOrder(userID int) {
	var gameID int
	var quantity int

	fmt.Print("Input Game ID: ")
	fmt.Scan(&gameID)

	fmt.Print("Input Quantity: ")
	fmt.Scan(&quantity)

	err := h.OrderRepo.CreateOrder(userID, gameID, quantity)

	if err != nil {
		fmt.Println("Create Order Failed:", err)
		return
	}

	fmt.Println("Create Order Success!")
}
