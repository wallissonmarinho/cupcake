package domain

import "time"

type UserRequest struct {
	ID         string    `json:"-" db:"id"`
	Password   string    `json:"password" db:"password"`
	Email      string    `json:"email" db:"email"`
	IsActive   bool      `json:"is_active" db:"is_active"`
	IsStaff    bool      `json:"is_staff" db:"is_staff"`
	FirstLogin time.Time `json:"first_login" db:"first_login"`
	LastLogin  time.Time `json:"last_login" db:"last_login"`
}

type UserProfileRequest struct {
	ID         string      `json:"-" db:"id"`
	FirstName  string      `json:"first_name" db:"first_name"`
	LastName   string      `json:"last_name" db:"last_name"`
	Phone      string      `json:"phone" db:"phone"`
	Address    string      `json:"address" db:"address"`
	City       string      `json:"city" db:"city"`
	State      string      `json:"state" db:"state"`
	PostalCode string      `json:"postal_code" db:"postal_code"`
	UserID     string      `json:"user_id" db:"user_id"`
	User       UserRequest `json:"user" db:"-"`
}
