package service

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"information-server/config"
	"information-server/helpers"
	"information-server/models"
)

type anatomicRegionHandler struct {
}

func (handler *anatomicRegionHandler) GetAnatomicRegion(c *gin.Context) {
	var anatomicRegion models.AnatomicRegion
	medicalSpecialtyId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("anatomicRegions").FindId(bson.ObjectIdHex(medicalSpecialtyId)).One(&anatomicRegion)
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(200, anatomicRegion)

}

func (handler *anatomicRegionHandler) CreateAnatomicRegion(c *gin.Context) {
	var anatomicRegion models.AnatomicRegion
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "body": nil})
	}
	err := c.Bind(&anatomicRegion)
	if err != nil {
		c.JSON(400, helpers.NewError("problem decoding body"))
		return
	}
	session := mongo.Session.Clone()
	err = session.DB(mongo.Database).C("anatomicRegions").Insert(&anatomicRegion)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, anatomicRegion)
}

func (handler *anatomicRegionHandler) UpdateAnatomicRegion(c *gin.Context) {

}

func (handler *anatomicRegionHandler) DeleteAnatomicRegion(c *gin.Context) {
	anatomicRegionId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("anatomicRegions").RemoveId(bson.ObjectIdHex(anatomicRegionId))
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(204, gin.H{})
}

func (handler *anatomicRegionHandler) GetAnatomicRegions(c *gin.Context) {
	anatomicRegionId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("anatomicRegions").RemoveId(bson.ObjectIdHex(anatomicRegionId))
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(204, gin.H{})
}
