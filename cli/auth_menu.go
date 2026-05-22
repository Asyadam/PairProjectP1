package cli

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/handler"
)

func AuthMenu(
	authHandler *handler.AuthHandler,
	gameHandler *handler.GameHandler,
	categoryHandler *handler.CategoryHandler,
	orderHandler *handler.OrderHandler,
) {
	var choice int

	for {
		fmt.Println("==== AUTH MENU ====")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Choose Menu: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			authHandler.Register()

		case 2:
			user, success := authHandler.Login()

			if success {
				MainMenu(gameHandler, categoryHandler, orderHandler, user.ID)
			}

		case 3:
			fmt.Println("Good Bye!")
			return

		default:
			fmt.Println("Invalid Menu")
		}

		fmt.Println()
	}
}
