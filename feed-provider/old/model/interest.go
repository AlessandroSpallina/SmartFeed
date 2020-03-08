package model

type Interest struct {
	User string              `json:"-"`
	Tag  string              `json:"tag"`
	Args map[string][]string `json:"args"`
}
