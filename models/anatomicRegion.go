package models

import "gopkg.in/mgo.v2/bson"

type AnatomicRegion struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name"`
}

func (c *AnatomicRegion) ToString() string {
	return "Name:" + c.Name
}
