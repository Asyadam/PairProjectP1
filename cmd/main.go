package main

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/cli"
	"github.com/Asyadam/PairProjectP1/db"
	"github.com/Asyadam/PairProjectP1/handler"
	"github.com/Asyadam/PairProjectP1/repository"
)

func main() {

	// connect database
	database, err := db.InitDB()

	if err != nil {
		fmt.Println("Failed connect database")
		return
	}

	// repository
	authRepo := repository.NewAuthRepository(database)

	// handler
	authHandler := handler.NewAuthHandler(authRepo)

	// run auth menu
	cli.AuthMenu(authHandler)
}
