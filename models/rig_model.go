package models

type Rig struct {
	BaseModel
	ChronoRig	map[string]interface{}	`json:"chrono_rig"`
	Patterns	[]Pattern        	`json:"patterns"`
	Leader		Leader        		`json:"leader"`
	Line 		Line        		`json:"line"`
}



