package cli

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/handler"
)

func MainMenu(
	gameHandler *handler.GameHandler,
	categoryHandler *handler.CategoryHandler,
	orderHandler *handler.OrderHandler,
	userID int,
) {
	var choice int

	for {
		fmt.Println("==== MAIN MENU ====")
		fmt.Println("1. Game Menu")
		fmt.Println("2. Category Menu")
		fmt.Println("3. Order Menu")
		fmt.Println("4. Logout")
		fmt.Print("Choose Menu: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			GameMenu(gameHandler)

		case 2:
			CategoryMenu(categoryHandler)

		case 3:
			OrderMenu(orderHandler, userID)

		case 4:
			fmt.Println("Logout Success!")
			return

		default:
			fmt.Println("Invalid Menu")
		}

		fmt.Println()
	}
}
