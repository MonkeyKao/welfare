package routes

import (
	"walfare/api"

	"github.com/gin-gonic/gin"
)

func WelfareRoutes(router *gin.Engine) {
	WelfareRoutes := router.Group("/welfare")
	{
		WelfareRoutes.GET("", api.GetWelfareHandler)
		WelfareRoutes.GET("/:id", api.GetWelfareByID)
	}
}
