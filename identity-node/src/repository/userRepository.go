package repository

// UserRepository - il controller tocca i repository, NON I MODEL
/*type UserRepository interface {
	Find() (model.User, error)
	Save(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(user model.User) error
}*/

import (
	"errors"
	"identity-node/src/model"
	"strings"
)

var userList = []model.User{
	model.User{Username: "user1", Password: "pass1"},
	model.User{Username: "user2", Password: "pass2"},
	model.User{Username: "user3", Password: "pass3"},
}

var sessionList = []model.Session{}

// SaveSession - Save user session
func SaveSession(user, token string) error {
	sessionList = append(sessionList, model.Session{User: user, Token: token})
	// @findme : una vera sessione dovrebbe scadere, da risistemare quando si mette il db

	return nil
}

// FindUserBySession - Return User by session token
func FindUserBySession(token string) (*model.User, error) {
	for _, s := range sessionList {
		if s.Token == token {
			return FindUser(s.User)
		}
	}
	return nil, errors.New("session not found")
}

// FindUser - Return User by username
func FindUser(username string) (*model.User, error) {
	for _, u := range userList {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

// IsValidUser - Check if the username and password combination is valid
func IsValidUser(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

// SaveUser - Save User
func SaveUser(u model.User) (*model.User, error) {
	if strings.TrimSpace(u.Password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if !isUsernameAvailable(u.Username) {
		return nil, errors.New("the username isn't available")
	}

	userList = append(userList, u)

	return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
