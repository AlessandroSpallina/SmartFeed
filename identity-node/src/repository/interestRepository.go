package repository

import (
	"identity-node/src/model"
)

// ListInterestsByUser -
func ListInterestsByUser(user string) []model.Interest {
	toReturn := []model.Interest{}

	for _, i := range interestList {
		if i.User == user {
			toReturn = append(toReturn, i)
		}
	}

	return toReturn
}

// SaveInterest - salva l'interesse dell'utente su "db"
func SaveInterest(i model.Interest) (*model.Interest, error) {
	_, err := FindUser(i.User)

	if err != nil {
		return nil, err
	}

	interestList = append(interestList, i)

	return &i, nil
}
