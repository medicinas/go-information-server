package service

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"information-server/config"
	"information-server/helpers"
	"information-server/models"
)

type countryResource struct {
}

func (handler *countryResource) GetCountries(c *gin.Context) {
	var countries []models.Country
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("countries").Find(bson.M{}).All(&countries)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, countries)
}

func (handler *countryResource) GetCountry(c *gin.Context) {
	var country models.Country
	countryId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("provinces").FindId(bson.ObjectIdHex(countryId)).One(&country)
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(200, country)

}

func (handler *countryResource) CreateCountry(c *gin.Context) {
	var province models.Province
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "body": nil})
	}
	err := c.Bind(&province)
	if err != nil {
		c.JSON(400, helpers.NewError("problem decoding body"))
		return
	}
	session := mongo.Session.Clone()
	err = session.DB(mongo.Database).C("provinces").Insert(&province)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, province)
}

func (handler *countryResource) UpdateCountry(c *gin.Context) {

}

func (handler *countryResource) DeleteCountry(c *gin.Context) {
	provinceId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("provinces").RemoveId(bson.ObjectIdHex(provinceId))
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(204, gin.H{})
}
