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
	medicineResource := medicineHandler{}
	provinceResource := provinceHandler{}
	cityResource := cityHandler{}
	anatomicRegionResource := anatomicRegionHandler{}

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(MiddleDB(&mongoDB))
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/medical-specialty", medicalSpecialityResource.GetMedicalSpecialties)
		apiv1.GET("/medical-specialty/:id", medicalSpecialityResource.GetMedicalSpecialty)
		apiv1.POST("/medical-specialty", medicalSpecialityResource.CreateMedicalSpecialty)
		apiv1.PUT("/medical-specialty/", medicalSpecialityResource.UpdateMedicalSpecialty)
		apiv1.DELETE("/medical-specialty/:id", medicalSpecialityResource.DeleteMedicalSpecialty)

		apiv1.GET("/cities", cityResource.GetCities)
		apiv1.GET("/cities/:id", cityResource.GetCity)
		apiv1.POST("/cities", cityResource.CreateCity)
		apiv1.PUT("/cities", cityResource.UpdateCity)
		apiv1.DELETE("/cities/:id", cityResource.DeleteCity)

		apiv1.GET("/provinces", provinceResource.GetProvinces)
		apiv1.GET("/provinces/:id", provinceResource.GetProvince)
		apiv1.POST("/provinces", provinceResource.CreateProvince)
		apiv1.PUT("/provinces", provinceResource.UpdateProvince)
		apiv1.DELETE("/provinces/:id", provinceResource.DeleteProvince)

		apiv1.GET("/medicines", medicineResource.GetMedicines)
		apiv1.GET("/medicines/:id", medicineResource.GetMedicine)
		apiv1.POST("/medicines", medicineResource.CreateMedicine)
		apiv1.PUT("/medicines", medicineResource.UpdateMedicine)
		apiv1.DELETE("/medicines/:id", medicineResource.DeleteMedicine)

		apiv1.GET("/anatomic-region", anatomicRegionResource.GetAnatomicRegions)
		apiv1.GET("/anatomic-region/:id", anatomicRegionResource.GetAnatomicRegion)
		apiv1.POST("/anatomic-region", anatomicRegionResource.CreateAnatomicRegion)
		apiv1.PUT("/anatomic-region", anatomicRegionResource.UpdateAnatomicRegion)
		apiv1.DELETE("/anatomic-region/:id", anatomicRegionResource.DeleteAnatomicRegion)

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
