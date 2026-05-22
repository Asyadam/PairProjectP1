package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Asyadam/PairProjectP1/handler"
)

func AuthMenu(
	authHandler *handler.AuthHandler,
	gameHandler *handler.GameHandler,
	categoryHandler *handler.CategoryHandler,
	orderHandler *handler.OrderHandler,
	reportHandler *handler.ReportHandler,
	profileHandler *handler.ProfileHandler,
) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("==== AUTH MENU ====")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Choose Menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid Menu")
			fmt.Println()
			continue
		}

		switch choice {
		case 1:
			authHandler.Register(reader)

		case 2:
			user, success := authHandler.Login(reader)

			if success {
				MainMenu(
					gameHandler,
					categoryHandler,
					orderHandler,
					reportHandler,
					profileHandler,
					reader,
					user.ID,
				)
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
