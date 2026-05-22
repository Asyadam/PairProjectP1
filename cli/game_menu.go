package cli

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/handler"
)

func GameMenu(gameHandler *handler.GameHandler) {
	var choice int

	for {
		fmt.Println("==== GAME MENU ====")
		fmt.Println("1. Add Game")
		fmt.Println("2. View Games")
		fmt.Println("3. Update Game")
		fmt.Println("4. Delete Game")
		fmt.Println("5. Back")
		fmt.Print("Choose Menu: ")

		fmt.Scan(&choice)
		fmt.Scanln()

		switch choice {
		case 1:
			gameHandler.CreateGame()

		case 2:
			gameHandler.ViewGames()

		case 3:
			gameHandler.UpdateGame()

		case 4:
			gameHandler.DeleteGame()

		case 5:
			return

		default:
			fmt.Println("Invalid Menu")
		}

		fmt.Println()
	}
}
