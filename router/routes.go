package router

import (
	docs "github.com/IgorAfonso/GoOportunitiesAPI/docs"
	"github.com/IgorAfonso/GoOportunitiesAPI/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHendler)
		v1.POST("/opening", handler.CreateOpeningHendler)
		v1.DELETE("/opening", handler.DeleteOpeningHendler)
		v1.PUT("/opening", handler.UpdateOpeningHendler)
		v1.GET("/openings", handler.ListOpeningHendler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}