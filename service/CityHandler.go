package service

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"medicinas/information-server/config"
	"medicinas/information-server/helpers"
	"medicinas/information-server/models"
)

type cityHandler struct {
}

func (handler *cityHandler) GetCities(c *gin.Context) {
	var cities []models.City
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("cities").Find(bson.M{}).All(&cities)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, cities)
}

func (handler *cityHandler) GetCity(c *gin.Context) {
	var city models.City
	medicalSpecialtyId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("cities").FindId(bson.ObjectIdHex(medicalSpecialtyId)).One(&city)
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(200, city)

}

func (handler *cityHandler) CreateCity(c *gin.Context) {
	var city models.City
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "body": nil})
	}
	err := c.Bind(&city)
	if err != nil {
		c.JSON(400, helpers.NewError("problem decoding body"))
		return
	}
	session := mongo.Session.Clone()
	err = session.DB(mongo.Database).C("cities").Insert(&city)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, city)
}

func (handler *cityHandler) UpdateCity(c *gin.Context) {

}

func (handler *cityHandler) DeleteCity(c *gin.Context) {
	cityId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("cities").RemoveId(bson.ObjectIdHex(cityId))
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(204, gin.H{})
}
