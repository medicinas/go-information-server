package models

import "gopkg.in/mgo.v2/bson"

type City struct {
	Id           bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string        `json:"name"`
	PostalCode   string        `json:"postalCode"`
	CountryCode  string        `json:"countryCode"`
	ProvinceCode string        `json:"provinceCode"`
}

func (c *City) ToString() string {
	return "Name:" + c.Name + " Postal code:" + c.PostalCode
}
