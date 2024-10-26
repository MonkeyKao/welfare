package routes

import (
	"walfare/api"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	UserRoutes := router.Group("/users")
	{
		UserRoutes.POST("/doLogin", api.DoLoginHandler)
		UserRoutes.POST("/register", api.RegisterHandler)
		UserRoutes.GET("/verify/:id", api.VerifyEmailHandler)
		UserRoutes.Use(api.JwtAuthMiddleware()).GET("", api.GetUserByUserIDHandler)
		UserRoutes.Use(api.JwtAuthMiddleware()).POST("/updateuser", api.UpdateuserHandler)
	}
}
