package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/Asyadam/PairProjectP1/handler"
)

func CategoryMenu(
	categoryHandler *handler.CategoryHandler,
	reader *bufio.Reader,
) {
	for {
		fmt.Println("==== CATEGORY MENU ====")
		fmt.Println("1. Add Category")
		fmt.Println("2. View Categories")
		fmt.Println("3. Update Category")
		fmt.Println("4. Delete Category")
		fmt.Println("5. Back")
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
			categoryHandler.CreateCategory()

		case 2:
			categoryHandler.ViewCategories()

		case 3:
			categoryHandler.UpdateCategory()

		case 4:
			categoryHandler.DeleteCategory()

		case 5:
			return

		default:
			fmt.Println("Invalid Menu")
		}

		fmt.Println()
	}
}
