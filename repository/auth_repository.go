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

func (r *AuthRepository) CreateUserWithProfile(user entity.User, profile entity.Profile) error {
	tx, err := r.DB.Begin()

	if err != nil {
		return err
	}

	result, err := tx.Exec(`
		INSERT INTO users(email, password, role)
		VALUES (?, ?, ?)
	`, user.Email, user.Password, user.Role)

	if err != nil {
		tx.Rollback()
		return err
	}

	userID, err := result.LastInsertId()

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO profiles(user_id, full_name, phone, address)
		VALUES (?, ?, ?, ?)
	`, userID, profile.FullName, profile.Phone, profile.Address)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
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
