package repository

import (
	"database/sql"
	"fmt"

	"github.com/Asyadam/PairProjectP1/entity"
)

type GameRepository struct {
	DB *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{DB: db}
}

func (r *GameRepository) CreateGame(game entity.Game) error {

	query := `
		INSERT INTO games(title, price, stock, description, release_date)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		game.Title,
		game.Price,
		game.Stock,
		game.Description,
		game.ReleaseDate,
	)

	return err
}

func (r *GameRepository) GetAllGames() ([]entity.Game, error) {

	var games []entity.Game

	query := `
		SELECT id, title, price, stock, description, release_date
		FROM games
	`

	rows, err := r.DB.Query(query)

	if err != nil {
		return games, err
	}

	defer rows.Close()

	for rows.Next() {

		var game entity.Game
		var releaseDate []byte

		err := rows.Scan(
			&game.ID,
			&game.Title,
			&game.Price,
			&game.Stock,
			&game.Description,
			&releaseDate,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}

		game.ReleaseDate = string(releaseDate)

		games = append(games, game)
	}

	return games, nil
}

func (r *GameRepository) UpdateGame(game entity.Game) error {

	query := `
		UPDATE games
		SET title = ?, price = ?, stock = ?, description = ?, release_date = ?
		WHERE id = ?
	`

	_, err := r.DB.Exec(
		query,
		game.Title,
		game.Price,
		game.Stock,
		game.Description,
		game.ReleaseDate,
		game.ID,
	)

	return err
}

func (r *GameRepository) DeleteGame(id int) error {

	query := `
		DELETE FROM games
		WHERE id = ?
	`

	_, err := r.DB.Exec(query, id)

	return err
}
