package cli

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/handler"
)

func ReportMenu(reportHandler *handler.ReportHandler) {
	var choice int

	for {
		fmt.Println("==== REPORT MENU ====")
		fmt.Println("1. User Report")
		fmt.Println("2. Sales Report")
		fmt.Println("3. Stock Report")
		fmt.Println("4. Back")
		fmt.Print("Choose Menu: ")

		fmt.Scan(&choice)

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
