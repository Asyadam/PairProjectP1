package repository

import (
	"database/sql"

	"github.com/Asyadam/PairProjectP1/entity"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) CreateUser(user entity.User) error {

	query := `
		INSERT INTO users(email, password, role)
		VALUES (?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		user.Email,
		user.Password,
		user.Role,
	)

	return err
}

func (r *AuthRepository) FindUserByEmail(email string) (entity.User, error) {

	var user entity.User

	query := `
		SELECT id, email, password, role, created_at
		FROM users
		WHERE email = ?
	`

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	return user, err
}
