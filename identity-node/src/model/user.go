package model

import "time"

// Session - sessioni di autenticazione attive per questo utente
type Session struct {
	User  string
	Token string
}

// User -> Utente-Turista
type User struct {
	Username    string    `json:"username"`
	Password    string    `json:"-"`
	DateOfBirth time.Time `json:"date-of-birth"`
	Gender      string    `json:"gender"`
	Phone       string    `json:"phone"`
	Email       string    `string:"email"`
}
