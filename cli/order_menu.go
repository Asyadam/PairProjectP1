package cli

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/handler"
)

func OrderMenu(orderHandler *handler.OrderHandler, userID int) {
	var choice int

	for {
		fmt.Println("==== ORDER MENU ====")
		fmt.Println("1. Create Order")
		fmt.Println("2. Back")
		fmt.Print("Choose Menu: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			orderHandler.CreateOrder(userID)

		case 2:
			return

		default:
			fmt.Println("Invalid Menu")
		}

		fmt.Println()
	}
}
