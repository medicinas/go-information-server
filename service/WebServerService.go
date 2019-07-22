package service

import (
	"github.com/gin-gonic/gin"
	"medicinas/information-server/config"
)

type WebServerService struct {
}

func (s *WebServerService) Run(port string) {

	mongoDB := config.MongoDB{}
	mongoDB.SetDefault()

	medicalSpecialityResource := medicalSpecialtyHandler{}

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(MiddleDB(&mongoDB))
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/medical-specialty", medicalSpecialityResource.GetMedicalSpecialties)
		apiv1.GET("/medical-specialty/:id", medicalSpecialityResource.GetMedicalSpecialty)
		apiv1.POST("/medical-specialty", medicalSpecialityResource.CreateMedicalSpecialty)
		apiv1.PUT("/medical-specialty/:id", medicalSpecialityResource.UpdateMedicalSpecialty)
		apiv1.DELETE("/medical-specialty/:id", medicalSpecialityResource.DeleteMedicalSpecialty)
	}

	router.StaticFile("/", "./public/index.html")
	router.Static("/public/", "./public/")
	_ = router.Run(":" + port)
}

func MiddleDB(mongo *config.MongoDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := mongo.SetSession()
		if err != nil {
			c.Abort()
		} else {
			c.Set("mongo", mongo)
			c.Next()
		}
	}
}
