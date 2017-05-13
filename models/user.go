package models

import (
	"database/sql"
)

// UserModel contains db and performs actions
type UserModel struct {
	Db *sql.DB
}

// User to be passed around to actions / json encoded in responses
type User struct {
	ID                   string `json:"id"`
	DisplayName          string `json:"display_name"`
	Email                string `json:"email"`
	Password             string `json:"password,omitempty"`
	PasswordConfirmation string `json:"password_confirmation,omitempty"`
}

// Create takes in a User and saves it to the database, returning the saved user
func (um *UserModel) Create(u *User) (*User, error) {
	var newUser User
	// bcrypt here for password storage
	passwordDigest := u.Password
	query := `
		INSERT INTO users
		(display_name, email, password_digest)
		VALUES ($1,$2, $3)
		RETURNING id
	`
	err := um.Db.QueryRow(query, u.DisplayName, u.Email, passwordDigest).Scan(&newUser.ID)
	if err != nil {
		return nil, err
	}
	return um.Get(&newUser)
}

// GetAll retreives all Users from the database and returns them in a list of Users
func (um *UserModel) GetAll() ([]*User, error) {
	query := `
		SELECT id, display_name, email
		FROM users
	`
	rows, err := um.Db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	users := []*User{}
	for rows.Next() {
		newUser := User{}
		err := rows.Scan(&newUser.ID, &newUser.DisplayName, &newUser.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &newUser)
	}
	return users, nil
}

// Get gets a specific user
func (um *UserModel) Get(u *User) (*User, error) {
	query := `
		SELECT id, display_name, email
		FROM users
		WHERE
		id = $1
	`
	row := um.Db.QueryRow(query, u.ID)
	var user User
	err := row.Scan(&user.ID, &user.DisplayName, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update takes in a user and updates it with new information
func (um *UserModel) Update(u *User) (*User, error) {
	query := `
		UPDATE users
		SET
		display_name = $2
		WHERE
		id = $1
	`
	_, err := um.Db.Exec(query, u.ID, u.DisplayName, &u.Email)
	if err != nil {
		return nil, err
	}
	return um.Get(u)
}
