package model

import "time"

// User -> Utente-Turista
type User struct {
	Username    string    `json:"username"`
	Password    string    `json:"-"`
	DateOfBirth time.Time `json:"date-of-birth"`
	Gender      rune      `json:"gender"`
	Phone       string    `json:"phone"`
	Email       string    `string:"email"`
}
