package repository

import (
	"../model"
)

func SaveInterest(i model.Interest) (*model.Interest, error) {
	_, err := FindUser(i.User)
	if err != nil {
		return nil, err
	}
	interestList = append(interestList, i)
	return &i, nil
}

func ListInterestByUser(user string) []model.Interest {
	toReturn := []model.Interest{}

	for _, i := range interestList {
		if i.User == user {
			toReturn = append(toReturn, i)
		}
	}

	return toReturn
}
