package router

import (
	"github.com/IgorAfonso/GoOportunitiesAPI/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHendler)
		v1.POST("/opening", handler.CreateOpeningHendler)
		v1.DELETE("/opening", handler.DeleteOpeningHendler)
		v1.PUT("/opening", handler.UpdateOpeningHendler)
		v1.GET("/openings", handler.ListOpeningHendler)
	}
}