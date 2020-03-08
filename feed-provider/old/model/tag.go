package model

type Tag struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}
