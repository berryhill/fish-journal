package models

type Line struct {
	BaseModel
	Name 		string          `json:"name"`
	Brand		string        	`json:"brand"`
	Kind  		string 		`json:"kind"`
}

