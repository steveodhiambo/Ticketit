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
	Username  string    `json:"username"`
	FirstName string    `json:"first"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
