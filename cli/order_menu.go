package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/Asyadam/PairProjectP1/handler"
)

func OrderMenu(
	orderHandler *handler.OrderHandler,
	reader *bufio.Reader,
	userID int,
) {
	for {
		fmt.Println("==== ORDER MENU ====")
		fmt.Println("1. Create Order")
		fmt.Println("2. Back")
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
			orderHandler.CreateOrder(userID)

		case 2:
			return

		default:
			fmt.Println("Invalid Menu")
		}

		fmt.Println()
	}
}
