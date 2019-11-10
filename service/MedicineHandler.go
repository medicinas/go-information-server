package service

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"medicinas/information-server/config"
	"medicinas/information-server/helpers"
	"medicinas/information-server/models"
)

type medicineHandler struct {
}

func (handler *medicineHandler) GetMedicines(c *gin.Context) {
	var medicines []models.Medicine
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("medicines").Find(bson.M{}).All(&medicines)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, medicines)
}

func (handler *medicineHandler) GetMedicine(c *gin.Context) {
	var medicine models.Medicine
	medicalSpecialtyId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("medicines").FindId(bson.ObjectIdHex(medicalSpecialtyId)).One(&medicine)
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medical specialty", "body": nil})
		return
	}
	c.JSON(200, medicine)

}

func (handler *medicineHandler) CreateMedicine(c *gin.Context) {
	var medicine models.Medicine
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "body": nil})
	}
	err := c.Bind(&medicine)
	if err != nil {
		c.JSON(400, helpers.NewError("problem decoding body"))
		return
	}
	session := mongo.Session.Clone()
	err = session.DB(mongo.Database).C("medicines").Insert(&medicine)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, medicine)
}

func (handler *medicineHandler) UpdateMedicine(c *gin.Context) {

}

func (handler *medicineHandler) DeleteMedicine(c *gin.Context) {
	medicineId := c.Param("id")
	mongo, ok := c.Keys["mongo"].(*config.MongoDB)
	if !ok {
		c.JSON(500, gin.H{"message": "can't reach db", "body": nil})
		return
	}
	session := mongo.Session.Clone()
	defer session.Close()
	err := session.DB(mongo.Database).C("medicines").RemoveId(bson.ObjectIdHex(medicineId))
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't create a medicine", "body": nil})
		return
	}
	c.JSON(204, gin.H{})
}
