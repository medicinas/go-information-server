package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type WebServerService struct {
}

func (s *WebServerService) Run(port string) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(viper.GetString("mongo_db"))
	client, _ := mongo.Connect(ctx, clientOptions)

	medicalSpecialityResource := medicalSpecialtyHandler{client: client}

	r := gin.Default()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/medical-specialty", medicalSpecialityResource.GetMedicalSpecialties)
		apiv1.GET("/medical-specialty/:id", medicalSpecialityResource.GetMedicalSpecialty)
		apiv1.POST("/medical-specialty", medicalSpecialityResource.CreateMedicalSpecialty)
		apiv1.PUT("/medical-specialty/:id", medicalSpecialityResource.UpdateMedicalSpecialty)
		apiv1.DELETE("/medical-specialty/:id", medicalSpecialityResource.DeleteMedicalSpecialty)
	}

	r.StaticFile("/", "./public/index.html")
	r.Static("/public/", "./public/")

	r.Run(":" + port)
}
