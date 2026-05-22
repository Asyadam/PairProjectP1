package repository

import (
	"database/sql"

	"github.com/Asyadam/PairProjectP1/entity"
)

type GameRepository struct {
	DB *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{DB: db}
}

func (r *GameRepository) CreateGame(game entity.Game, categoryIDs []int) error {
	tx, err := r.DB.Begin()

	if err != nil {
		return err
	}

	result, err := tx.Exec(`
		INSERT INTO games(title, price, stock, description, release_date)
		VALUES (?, ?, ?, ?, ?)
	`,
		game.Title,
		game.Price,
		game.Stock,
		game.Description,
		game.ReleaseDate,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	gameID, err := result.LastInsertId()

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, categoryID := range categoryIDs {
		_, err = tx.Exec(`
			INSERT INTO game_categories(game_id, category_id)
			VALUES (?, ?)
		`, gameID, categoryID)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (r *GameRepository) GetAllGames() ([]entity.Game, error) {
	var games []entity.Game

	query := `
		SELECT 
			g.id,
			g.title,
			g.price,
			g.stock,
			COALESCE(g.description, ''),
			COALESCE(DATE_FORMAT(g.release_date, '%Y-%m-%d'), ''),
			COALESCE(GROUP_CONCAT(c.category_name SEPARATOR ', '), '-')
		FROM games g
		LEFT JOIN game_categories gc ON g.id = gc.game_id
		LEFT JOIN categories c ON gc.category_id = c.id
		GROUP BY 
			g.id,
			g.title,
			g.price,
			g.stock,
			g.description,
			g.release_date
	`

	rows, err := r.DB.Query(query)

	if err != nil {
		return games, err
	}

	defer rows.Close()

	for rows.Next() {
		var game entity.Game

		err := rows.Scan(
			&game.ID,
			&game.Title,
			&game.Price,
			&game.Stock,
			&game.Description,
			&game.ReleaseDate,
			&game.Categories,
		)

		if err != nil {
			return games, err
		}

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
