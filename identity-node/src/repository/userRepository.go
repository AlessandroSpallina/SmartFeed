package repository

import (
	"errors"
	"identity-node/src/model"
	"strings"
)

// SaveSession - Save user session on "db"
func SaveSession(user, token string) error {
	_, err := FindUser(user)

	if err != nil {
		return err
	}

	sessionList = append(sessionList, model.Session{User: user, Token: token})
	// @findme : una vera sessione dovrebbe scadere, da risistemare quando si mette il db

	return nil
}

// FindUserBySession - Return user from "db" by session token
func FindUserBySession(token string) (*model.User, error) {
	for _, s := range sessionList {
		if s.Token == token {
			return FindUser(s.User)
		}
	}
	return nil, errors.New("session not found")
}

// FindUser - Return user from "db" by username
func FindUser(username string) (*model.User, error) {
	for _, u := range userList {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

// IsValidUser - Check if the username and password combination is valid (ergo combination present in "db")
func IsValidUser(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

// SaveUser - Save user in "db"
// @findme : la validation dovrebbe essere in un controller o middleware!
// ma questa è validation? questo è controllo consistenza db, quindi forse può stare qui
// o meglio, diciamo che il controllo di password non vuoto è validation, la verifica di username disponibile è consisenza db
// quindi dovresti spostare la validation password empty solamente, quando hai tempo e con calma :D
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
