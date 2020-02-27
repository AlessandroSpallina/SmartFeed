package model

// Interest - Interest{Name:"meteo", Args:{"city":["catania", "giarre"]}}
type Interest struct {
	User string
	Name string
	Args map[string][]string
}
