package routes

import (
	"walfare/api"

	"github.com/gin-gonic/gin"
)

func FamilyRoutes(c *gin.Engine) {
	FamilyRoutes := c.Group("/family", api.JwtAuthMiddleware())
	{
		FamilyRoutes.POST("/:name", api.CreateFamilyHandler)
		FamilyRoutes.GET("", api.GetAllFamilyByUserIDHandler)
		FamilyRoutes.GET("/bind/:familyId", api.GetFmailyCodeHandler)
		FamilyRoutes.POST("/join/:code", api.JoinFmailyHandler)
		FamilyRoutes.DELETE("/:id", api.DeleteFamilyHandler)
		FamilyRoutes.DELETE("/leave/:id", api.LeaveFmailyHandler)
	}
}
