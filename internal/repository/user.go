package repository

import (
	"errors"

	"github.com/with-insomnia/profile/internal/model"
)

func (r *Repository) CheckUser(user model.Credintails) error {
	var u model.User
	err := r.db.QueryRow("SELECT * FROM users WHERE username = $1", user.Username).Scan(&u.UserId, &u.Username, &u.Password)
	if err != nil {
		return err
	}
	if user.Password != u.Password {
		return errors.New("password doesn't match")
	}

	return nil
}
