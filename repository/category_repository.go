package repository

import (
	"database/sql"

	"github.com/Asyadam/PairProjectP1/entity"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) CreateCategory(category entity.Category) error {
	query := `
		INSERT INTO categories(category_name)
		VALUES (?)
	`

	_, err := r.DB.Exec(query, category.CategoryName)

	return err
}

func (r *CategoryRepository) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category

	query := `
		SELECT id, category_name
		FROM categories
	`

	rows, err := r.DB.Query(query)

	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		var category entity.Category

		err := rows.Scan(
			&category.ID,
			&category.CategoryName,
		)

		if err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) UpdateCategory(category entity.Category) error {
	query := `
		UPDATE categories
		SET category_name = ?
		WHERE id = ?
	`

	_, err := r.DB.Exec(
		query,
		category.CategoryName,
		category.ID,
	)

	return err
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	query := `
		DELETE FROM categories
		WHERE id = ?
	`

	_, err := r.DB.Exec(query, id)

	return err
}
