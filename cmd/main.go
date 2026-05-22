package main

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/cli"
	"github.com/Asyadam/PairProjectP1/db"
	"github.com/Asyadam/PairProjectP1/handler"
	"github.com/Asyadam/PairProjectP1/repository"
)

func main() {
	database, err := db.InitDB()

	if err != nil {
		fmt.Println("Failed connect database")
		return
	}

	// auth
	authRepo := repository.NewAuthRepository(database)
	authHandler := handler.NewAuthHandler(authRepo)

	// game
	gameRepo := repository.NewGameRepository(database)
	gameHandler := handler.NewGameHandler(gameRepo)

	// category
	categoryRepo := repository.NewCategoryRepository(database)
	categoryHandler := handler.NewCategoryHandler(categoryRepo)

	// run app
	cli.AuthMenu(authHandler, gameHandler, categoryHandler)
}
