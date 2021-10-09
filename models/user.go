package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

type Post struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Caption         string        `json:"caption" bson:"caption"`
	Imageurl        string        `json:"imageurl" bson:"imageurl"`
	PostedTimeStamp time.Time     `json:"PostedTimeStamp" bson:"ts"`
}
