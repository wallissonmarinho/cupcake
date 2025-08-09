package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/wallissonmarinho/cupcake/internal/domain"
)

type UserRepository interface {
	GetUserByID(id string) (domain.UserRequest, error)
	CreateUser(user domain.UserRequest) error
	UpdateUser(user domain.UserRequest) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByID(id string) (domain.UserRequest, error) {
	user := domain.UserRequest{}

	query := `
	SELECT 
		id, 
		password,
		email, 
		is_active, 
		is_staff, 
		first_login, 
		last_login 
	FROM users 
	WHERE id = $1`
	err := r.db.Get(&user, query, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user domain.UserRequest) error {
	query := `
	INSERT INTO users (id, email, password, is_active, is_staff) 
	VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query, user.ID, user.Email, user.Password, user.IsActive, user.IsStaff)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateUser(user domain.UserRequest) error {
	query := `
	UPDATE users 
	SET 
	email = $1, 
	password = $2, 
	is_active = $3, 
	is_staff = $4 
	WHERE id = $5`
	_, err := r.db.Exec(query, user.Email, user.Password, user.IsActive, user.IsStaff, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(id string) error {
	query := `
	DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
