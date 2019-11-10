package models

import "gopkg.in/mgo.v2/bson"

type Drug struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name"`
}

func (d *Drug) ToString() string {
	return "Name: " + d.Name
}
