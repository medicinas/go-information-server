package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"medicinas/information-server/helpers"
	"medicinas/information-server/models"
	"strconv"
	"time"
)

type medicalSpecialtyHandler struct {
	client *mongo.Client
}

func (handler *medicalSpecialtyHandler) GetMedicalSpecialties(c *gin.Context) {
	var specialities []models.MedicalSpecialty
	collection := handler.client.Database("basic_information").Collection("medical_specialties")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, err)
		return
	}
	for cursor.Next(ctx) {
		var specialty models.MedicalSpecialty
		err = cursor.Decode(&specialty)
		if err != nil {
			c.JSON(500, err)
		}
		specialities = append(specialities, specialty)
	}
	c.JSON(200, specialities)
}

func (handler *medicalSpecialtyHandler) GetMedicalSpecialty(c *gin.Context) {

}

func (handler *medicalSpecialtyHandler) CreateMedicalSpecialty(c *gin.Context) {
	var medicalSpecialty models.MedicalSpecialty
	err := c.Bind(&medicalSpecialty)
	if err != nil {
		c.JSON(400, helpers.NewError("problem decoding body"))
		return
	}
	collection := handler.client.Database("basic_information").Collection("medical_specialties")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, medicalSpecialty)
	c.JSON(200, result)
}

func (handler *medicalSpecialtyHandler) UpdateMedicalSpecialty(c *gin.Context) {

}

func (handler *medicalSpecialtyHandler) DeleteMedicalSpecialty(c *gin.Context) {

}

func (p *medicalSpecialtyHandler) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return int32(id), nil
}
