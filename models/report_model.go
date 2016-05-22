package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Report struct {
	BaseModel
	TimeStarted		time.Time       		`json:"time_started"`
	TimeEnded		time.Time       	 	`json:"time_ended"`
	events 			[]interface{}   	 	`json:"events"`
	EventIds		[]bson.ObjectId        	 	`json:"eventsids"`
	NumEvents		int 		       	 	`json:"num_events"`
	ChronoConditions	map[time.Time]Conditions        `json:"chrono_conditions"`
}

func NewReport() *Report {
	r := new(Report)
	r.Id = bson.NewObjectId()
	return r
}

func (r *Report) StartFishing() {
	r.TimeStarted = time.Now()
}

func (r *Report) FinishFishing() {
	r.TimeEnded = time.Now()
}

func (r *Report) AddEvent(e *Event) {
	r.EventIds[r.NumEvents] = e.Id
}

func (r *Report) GetAllEvents() {
	// TODO implement
}











