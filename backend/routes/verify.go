package routes

import (
	"walfare/api"

	"github.com/gin-gonic/gin"
)

func VerifyRoutes(router *gin.Engine) {
	VerifyRoutes := router.Group("/verify")
	{
		VerifyRoutes.POST("", api.VerifyEmailHandler)
		VerifyRoutes.GET("", api.GetVerifyEmailHandler)
	}
}
