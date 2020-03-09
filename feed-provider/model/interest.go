package model

type Interest struct {
	Tag  string              `json:"tag"`
	Args map[string][]string `json:"args"`
}
