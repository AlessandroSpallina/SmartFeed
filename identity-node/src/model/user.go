package model

// User -> Utente-Turista
type User struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	//DateOfBirth time.Time `json:"date-of-birth" time_format:"1994-04-30"`
	DateOfBirth string `json:"date-of-birth,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `string:"email,omitempty"`
}
