package repository

import (
	"errors"

	"../model"
)

func FindUser(username string) (*model.User, error) {
	for _, u := range userList {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
