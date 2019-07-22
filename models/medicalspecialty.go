package models

import "gopkg.in/mgo.v2/bson"

type MedicalSpecialty struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
}

func (a *MedicalSpecialty) ToString() string {
	return a.Name
}
