package routes

import (
	"walfare/api"

	"github.com/gin-gonic/gin"
)

func FavoriteRoutes(router *gin.Engine) {
	FavoriteRoutes := router.Group("/favorite", api.JwtAuthMiddleware())
	{
		FavoriteRoutes.POST("/:id", api.AddFavorite)
		FavoriteRoutes.GET("", api.GetFavorites)
		FavoriteRoutes.DELETE("/:id", api.DeleteFavorite)
	}
}
