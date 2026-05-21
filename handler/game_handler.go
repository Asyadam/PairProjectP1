package handler

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/entity"
	"github.com/Asyadam/PairProjectP1/repository"
)

type GameHandler struct {
	GameRepo *repository.GameRepository
}

func NewGameHandler(gameRepo *repository.GameRepository) *GameHandler {
	return &GameHandler{GameRepo: gameRepo}
}

func (h *GameHandler) CreateGame() {

	var game entity.Game

	fmt.Print("Input Title: ")
	fmt.Scan(&game.Title)

	fmt.Print("Input Price: ")
	fmt.Scan(&game.Price)

	fmt.Print("Input Stock: ")
	fmt.Scan(&game.Stock)

	fmt.Print("Input Description: ")
	fmt.Scan(&game.Description)

	fmt.Print("Input Release Date (YYYY-MM-DD): ")
	fmt.Scan(&game.ReleaseDate)

	err := h.GameRepo.CreateGame(game)

	if err != nil {
		fmt.Println("Create Game Failed:", err)
		return
	}

	fmt.Println("Create Game Success!")
}

func (h *GameHandler) ViewGames() {

	games, err := h.GameRepo.GetAllGames()

	if err != nil {
		fmt.Println("Get Games Failed:", err)
		return
	}

	fmt.Println("==== GAME LIST ====")

	for _, game := range games {

		fmt.Println("ID:", game.ID)
		fmt.Println("Title:", game.Title)
		fmt.Println("Price:", game.Price)
		fmt.Println("Stock:", game.Stock)
		fmt.Println("Description:", game.Description)
		fmt.Println("Release Date:", game.ReleaseDate)
		fmt.Println("====================")
	}
}

func (h *GameHandler) UpdateGame() {

	var game entity.Game

	fmt.Print("Input Game ID: ")
	fmt.Scan(&game.ID)

	fmt.Print("Input New Title: ")
	fmt.Scan(&game.Title)

	fmt.Print("Input New Price: ")
	fmt.Scan(&game.Price)

	fmt.Print("Input New Stock: ")
	fmt.Scan(&game.Stock)

	fmt.Print("Input New Description: ")
	fmt.Scan(&game.Description)

	fmt.Print("Input New Release Date (YYYY-MM-DD): ")
	fmt.Scan(&game.ReleaseDate)

	err := h.GameRepo.UpdateGame(game)

	if err != nil {
		fmt.Println("Update Game Failed:", err)
		return
	}

	fmt.Println("Update Game Success!")
}

func (h *GameHandler) DeleteGame() {

	var id int

	fmt.Print("Input Game ID: ")
	fmt.Scan(&id)

	err := h.GameRepo.DeleteGame(id)

	if err != nil {
		fmt.Println("Delete Game Failed:", err)
		return
	}

	fmt.Println("Delete Game Success!")
}
