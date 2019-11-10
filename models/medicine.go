package models

import "gopkg.in/mgo.v2/bson"

type Medicine struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name"`
	Code string        `json:"code"`
	Drug Drug          `json:"drugs"`
}

func (c *Medicine) ToString() string {
	return "Name:" + c.Name + " Code:" + c.Code
}
