package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/Asyadam/PairProjectP1/handler"
)

func ProfileMenu(
	profileHandler *handler.ProfileHandler,
	reader *bufio.Reader,
	userID int,
) {
	for {
		fmt.Println("==== PROFILE MENU ====")
		fmt.Println("1. View My Profile")
		fmt.Println("2. Update My Profile")
		fmt.Println("3. Back")
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
			profileHandler.ViewProfile(userID)

		case 2:
			profileHandler.UpdateProfile(reader, userID)

		case 3:
			return

		default:
			fmt.Println("Invalid Menu")
		}

		fmt.Println()
	}
}
