package models

import "time"

type Catch struct {
	BaseModel
	Event
	Rig
	Conditions
	Specie		Specie        	`json:"specie"`
	Length		int        	`json:"length"`
	Pattern		Pattern        	`json:"pattern"`
}

func NewCatch(specie Specie, length int) *Catch {
	c := new(Catch)
	c.Specie = specie
	c.Length = length
	c.Time = time.Now()
	return c
}




