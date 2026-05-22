package repository

import (
	"database/sql"

	"github.com/Asyadam/PairProjectP1/entity"
)

type ProfileRepository struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (r *ProfileRepository) GetProfileByUserID(userID int) (entity.Profile, error) {
	var profile entity.Profile

	query := `
		SELECT id, user_id, full_name, phone, address
		FROM profiles
		WHERE user_id = ?
	`

	err := r.DB.QueryRow(query, userID).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.FullName,
		&profile.Phone,
		&profile.Address,
	)

	return profile, err
}

func (r *ProfileRepository) UpdateProfile(profile entity.Profile) error {
	query := `
		UPDATE profiles
		SET full_name = ?, phone = ?, address = ?
		WHERE user_id = ?
	`

	_, err := r.DB.Exec(
		query,
		profile.FullName,
		profile.Phone,
		profile.Address,
		profile.UserID,
	)

	return err
}
