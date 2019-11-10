package models

import "gopkg.in/mgo.v2/bson"

type Province struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	CountryCode string        `json:"countryCode"`
}

func (c *Province) ToString() string {
	return "Name:" + c.Name
}
