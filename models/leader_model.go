package models

type Leader struct {
	BaseModel
	Length		float32        	`json:"length"`
	Taper 		string 		`json:"taper"`
	Kind		string        	`json:"kind"`
}

