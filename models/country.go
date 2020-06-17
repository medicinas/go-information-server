package models

import "gopkg.in/mgo.v2/bson"

type Country struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name"`
	Code string        `json:"code"`
}

func (c *Country) ToString() string {
	return "Name:" + c.Name + " Code:" + c.Code
}
