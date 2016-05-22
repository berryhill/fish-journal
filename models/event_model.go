package models

import "time"

type Event struct {
	BaseModel
	Name 		string        	`json:"name"`
	Coordinates	[]float32     	`json:"Coordinates"`
	Time 		time.Time 	`json:"time"`
}
