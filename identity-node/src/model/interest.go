package model

// Interest - Interest{User:"sk3la", Name:"meteo", Args:{"city":["catania", "giarre"]}}
type Interest struct {
	User string              `json:"-"`
	Tag  string              `json:"tag"`
	Args map[string][]string `json:"args"`
}
