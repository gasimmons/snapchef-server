package db

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"github.com/gasimmons/snapchef-server/auth"
)
import _ "github.com/gasimmons/snapchef-server/auth"

type User struct {
	ID           int    `db:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

func CreateUser(firstName string, lastName string, email string, rawPassword string) (int64, error) {
	hash, err := auth.HashPassword(rawPassword)
	if err != nil {
		return 0, err
	}

	res, err := DB.Exec(`
		INSERT INTO users (firstName, lastName, email, passwordHash)
		VALUES (?, ?, ?, ?)`,
		firstName, lastName, email, hash,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func GetUserByEmail(email string) (*User, error) {
	row := DB.QueryRow(`
	SELECT userId, firstName, lastName, email, passwordHash
	FROM users WHERE email=?`, email)

	var user User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
