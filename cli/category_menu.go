package cli

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/handler"
)

func CategoryMenu(categoryHandler *handler.CategoryHandler) {
	var choice int

	for {
		fmt.Println("==== CATEGORY MENU ====")
		fmt.Println("1. Add Category")
		fmt.Println("2. View Categories")
		fmt.Println("3. Update Category")
		fmt.Println("4. Delete Category")
		fmt.Println("5. Back")
		fmt.Print("Choose Menu: ")

		fmt.Scan(&choice)

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
