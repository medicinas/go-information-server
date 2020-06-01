package service

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"information-server/config"
	"information-server/helpers"
	"information-server/models"
)

type medicalSpecialtyHandler struct {
}

func (handler *medicalSpecialtyHandler) GetMedicalSpecialties(c *gin.Context) {
	var medicalSpecialties []models.MedicalSpecialty
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("medical_specialties").Find(bson.M{}).All(&medicalSpecialties)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, medicalSpecialties)
}

func (handler *medicalSpecialtyHandler) GetMedicalSpecialty(c *gin.Context) {
	var medicalSpecialty models.MedicalSpecialty
	medicalSpecialtyId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("medical_specialties").FindId(bson.ObjectIdHex(medicalSpecialtyId)).One(&medicalSpecialty)
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(200, medicalSpecialty)

}

func (handler *medicalSpecialtyHandler) CreateMedicalSpecialty(c *gin.Context) {
	var medicalSpecialty models.MedicalSpecialty
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "body": nil})
	}
	err := c.Bind(&medicalSpecialty)
	if err != nil {
		c.JSON(400, helpers.NewError("problem decoding body"))
		return
	}
	session := mongo.Session.Clone()
	err = session.DB(mongo.Database).C("medical_specialties").Insert(&medicalSpecialty)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, medicalSpecialty)
}

func (handler *medicalSpecialtyHandler) UpdateMedicalSpecialty(c *gin.Context) {

}

func (handler *medicalSpecialtyHandler) DeleteMedicalSpecialty(c *gin.Context) {
	medicalSpecialtyId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("medical_specialties").RemoveId(bson.ObjectIdHex(medicalSpecialtyId))
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(204, gin.H{})
}
