package model

// Tag - Tag{Name: "meteo", Args:["city", ]}
type Tag struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}
