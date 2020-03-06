package repository

import "../model"

var userList = []model.User{
	model.User{Username: "user1", Password: "pass1"},
	model.User{Username: "user2", Password: "pass2"},
	model.User{Username: "user3", Password: "pass3"},
}

var interestList = []model.Interest{}

var tagList = []model.Tag{
	model.Tag{Name: "monuments", Args: []string{"city"}},
	model.Tag{Name: "local-food", Args: []string{"city"}},
	model.Tag{Name: "museums", Args: []string{"city"}},
	model.Tag{Name: "trip", Args: []string{"city"}},
	model.Tag{Name: "teather", Args: []string{"city"}},
	model.Tag{Name: "local-events", Args: []string{"city"}},
	model.Tag{Name: "urban-transport", Args: []string{"city"}},
	model.Tag{Name: "suburban-transport", Args: []string{"from_city", "to_city"}},
	model.Tag{Name: "weather", Args: []string{"city"}},
	model.Tag{Name: "nightlife", Args: []string{"city"}},
}
