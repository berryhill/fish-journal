package models

type Pattern struct {
	BaseModel
	Name 		string     `json:"name"`
	Size 		int        `json:"size"`
}
