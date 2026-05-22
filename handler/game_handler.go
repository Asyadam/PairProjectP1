package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	var categoryIDs []int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input Title: ")
	title, _ := reader.ReadString('\n')
	game.Title = strings.TrimSpace(title)

	fmt.Print("Input Price: ")
	priceInput, _ := reader.ReadString('\n')
	priceInput = strings.TrimSpace(priceInput)

	price, err := strconv.ParseFloat(priceInput, 64)

	if err != nil {
		fmt.Println("Invalid price")
		return
	}

	game.Price = price

	fmt.Print("Input Stock: ")
	stockInput, _ := reader.ReadString('\n')
	stockInput = strings.TrimSpace(stockInput)

	stock, err := strconv.Atoi(stockInput)

	if err != nil {
		fmt.Println("Invalid stock")
		return
	}

	game.Stock = stock

	fmt.Print("Input Description: ")
	description, _ := reader.ReadString('\n')
	game.Description = strings.TrimSpace(description)

	fmt.Print("Input Release Date (YYYY-MM-DD): ")
	releaseDate, _ := reader.ReadString('\n')
	game.ReleaseDate = strings.TrimSpace(releaseDate)

	fmt.Print("Input Category IDs (contoh: 1,2,3): ")
	categoryInput, _ := reader.ReadString('\n')
	categoryInput = strings.TrimSpace(categoryInput)

	categoryParts := strings.Split(categoryInput, ",")

	for _, part := range categoryParts {
		part = strings.TrimSpace(part)

		categoryID, err := strconv.Atoi(part)

		if err != nil {
			fmt.Println("Invalid category ID:", part)
			return
		}

		categoryIDs = append(categoryIDs, categoryID)
	}

	err = h.GameRepo.CreateGame(game, categoryIDs)

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

	if len(games) == 0 {
		fmt.Println("No games found.")
		return
	}

	for _, game := range games {
		fmt.Println("ID:", game.ID)
		fmt.Println("Title:", game.Title)
		fmt.Println("Price:", game.Price)
		fmt.Println("Stock:", game.Stock)
		fmt.Println("Description:", game.Description)
		fmt.Println("Release Date:", game.ReleaseDate)
		fmt.Println("Categories:", game.Categories)
		fmt.Println("====================")
	}
}

func (h *GameHandler) UpdateGame() {
	var game entity.Game

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input Game ID: ")
	idInput, _ := reader.ReadString('\n')
	idInput = strings.TrimSpace(idInput)

	id, err := strconv.Atoi(idInput)

	if err != nil {
		fmt.Println("Invalid game ID")
		return
	}

	game.ID = id

	fmt.Print("Input New Title: ")
	title, _ := reader.ReadString('\n')
	game.Title = strings.TrimSpace(title)

	fmt.Print("Input New Price: ")
	priceInput, _ := reader.ReadString('\n')
	priceInput = strings.TrimSpace(priceInput)

	price, err := strconv.ParseFloat(priceInput, 64)

	if err != nil {
		fmt.Println("Invalid price")
		return
	}

	game.Price = price

	fmt.Print("Input New Stock: ")
	stockInput, _ := reader.ReadString('\n')
	stockInput = strings.TrimSpace(stockInput)

	stock, err := strconv.Atoi(stockInput)

	if err != nil {
		fmt.Println("Invalid stock")
		return
	}

	game.Stock = stock

	fmt.Print("Input New Description: ")
	description, _ := reader.ReadString('\n')
	game.Description = strings.TrimSpace(description)

	fmt.Print("Input New Release Date (YYYY-MM-DD): ")
	releaseDate, _ := reader.ReadString('\n')
	game.ReleaseDate = strings.TrimSpace(releaseDate)

	err = h.GameRepo.UpdateGame(game)

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
