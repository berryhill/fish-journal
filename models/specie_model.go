package models

type Specie struct {
	BaseModel
	Name 		string        `json:"name"`
	WikiLink 	string        `json:"wiki_link"`
}

func NewSpecie(name string) *Specie {
	s := new(Specie)
	s.Name = name
	return s
}