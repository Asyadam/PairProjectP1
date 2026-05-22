package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/Asyadam/PairProjectP1/handler"
)

func MainMenu(
	gameHandler *handler.GameHandler,
	categoryHandler *handler.CategoryHandler,
	orderHandler *handler.OrderHandler,
	reportHandler *handler.ReportHandler,
	profileHandler *handler.ProfileHandler,
	reader *bufio.Reader,
	userID int,
) {
	for {
		fmt.Println("==== MAIN MENU ====")
		fmt.Println("1. Game Menu")
		fmt.Println("2. Category Menu")
		fmt.Println("3. Order Menu")
		fmt.Println("4. Report Menu")
		fmt.Println("5. My Profile")
		fmt.Println("6. Logout")
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
			GameMenu(gameHandler)

		case 2:
			CategoryMenu(categoryHandler, reader)

		case 3:
			OrderMenu(orderHandler, reader, userID)

		case 4:
			ReportMenu(reportHandler, reader)

		case 5:
			ProfileMenu(profileHandler, reader, userID)

		case 6:
			fmt.Println("Logout Success!")
			return

		default:
			fmt.Println("Invalid Menu")
		}

		fmt.Println()
	}
}
