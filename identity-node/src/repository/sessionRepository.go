package repository

import (
	"identity-node/src/model"
)

// SaveSession - Save user session on "db"
func SaveSession(s model.Session) (*model.Session, error) {
	_, err := FindUser(s.User)

	if err != nil {
		return nil, err
	}

	sessionList = append(sessionList, s)
	// @findme : una vera sessione dovrebbe scadere, da risistemare quando si mette il db

	return &s, nil
}

/*func SaveSession(user, token string) error {
	_, err := FindUser(user)

	if err != nil {
		return err
	}

	sessionList = append(sessionList, model.Session{User: user, Token: token})
	// @findme : una vera sessione dovrebbe scadere, da risistemare quando si mette il db

	return nil
}*/
