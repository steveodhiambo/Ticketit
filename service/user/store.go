package user

import (
	"database/sql"
	"fmt"
	"github.com/steveodhiambo/ticket-it/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	// Pointer to user
	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	//User not found
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *Store) GetUserById(id int64) (*types.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) CreateUser(user types.User) error {
	//TODO implement me
	panic("implement me")
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}