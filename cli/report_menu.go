package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/Asyadam/PairProjectP1/handler"
)

func ReportMenu(
	reportHandler *handler.ReportHandler,
	reader *bufio.Reader,
) {
	for {
		fmt.Println("==== REPORT MENU ====")
		fmt.Println("1. User Report")
		fmt.Println("2. Sales Report")
		fmt.Println("3. Stock Report")
		fmt.Println("4. Back")
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
			reportHandler.UserReport()

		case 2:
			reportHandler.SalesReport()

		case 3:
			reportHandler.StockReport()

		case 4:
			return

		default:
			fmt.Println("Invalid Menu")
		}

		fmt.Println()
	}
}
