package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type BaseModel struct {
	Id		bson.ObjectId	`json:"id"`
	Timestamp	time.Time	`json:"timestamp"`
}


