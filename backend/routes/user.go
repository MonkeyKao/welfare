package routes

import (
	"walfare/api"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	UserRoutes := router.Group("/users")
	{
		// 不需要 JwtAuthMiddleware 的路由
		UserRoutes.POST("/doLogin", api.DoLoginHandler)
		UserRoutes.POST("", api.RegisterHandler)

		// 需要 JwtAuthMiddleware 的路由
		protectedRoutes := UserRoutes.Use(api.JwtAuthMiddleware())
		protectedRoutes.GET("", api.GetUserByUserIDHandler)
		protectedRoutes.PUT("", api.UpdateuserHandler)
		protectedRoutes.GET("/bind", api.GetBindQrCode)
	}
}
