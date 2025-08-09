package profile

import (
	"github.com/jmoiron/sqlx"
	"github.com/wallissonmarinho/cupcake/internal/domain"
)

type ProfileRepository interface {
	GetProfileByUserID(userID string) (domain.UserProfileRequest, error)
	CreateProfile(profile domain.UserProfileRequest) error
	UpdateProfile(profile domain.UserProfileRequest) error
	DeleteProfile(userID string) error
}

type profileRepository struct {
	db *sqlx.DB
}

func NewProfileRepository(db *sqlx.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

func (r *profileRepository) GetProfileByUserID(userID string) (domain.UserProfileRequest, error) {
	profile := domain.UserProfileRequest{}
	query := `
	SELECT 
		id, 
		first_name, 
		last_name, 
		phone, 
		address, 
		city, 
		state, 
		postal_code, 
		user_id 
	FROM user_profiles 
	WHERE user_id = $1`
	err := r.db.Get(&profile, query, userID)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (r *profileRepository) CreateProfile(profile domain.UserProfileRequest) error {
	query := `
	INSERT INTO user_profiles (id, first_name, last_name, phone, address, city, state, postal_code, user_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.db.Exec(query, profile.ID, profile.FirstName, profile.LastName, profile.Phone,
		profile.Address, profile.City, profile.State, profile.PostalCode, profile.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *profileRepository) UpdateProfile(profile domain.UserProfileRequest) error {
	query := `
	UPDATE user_profiles 
	SET first_name = $1, last_name = $2, phone = $3, address = $4, city = $5, state = $6, postal_code = $7 
	WHERE user_id = $8`

	_, err := r.db.Exec(query, profile.FirstName, profile.LastName, profile.Phone,
		profile.Address, profile.City, profile.State, profile.PostalCode, profile.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *profileRepository) DeleteProfile(userID string) error {
	query := `
	DELETE FROM user_profiles 
	WHERE user_id = $1`

	_, err := r.db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}
