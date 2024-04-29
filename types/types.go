package types

import "time"

// UserStore interface to allow testing
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int64) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"user_name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	Username  string `json:"user_name" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=32"`
}
