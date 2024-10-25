package routes

import(
	"github.com/gin-gonic/gin"
	"walfare/api"
)

func UserRoutes(router *gin.Engine){
	UserRoutes := router.Group("/users")
	{
		UserRoutes.POST("/doLogin",api.DoLoginHandler)
		UserRoutes.POST("/register",api.RegisterHandler)
		UserRoutes.Use(api.JwtAuthMiddleware()).POST("/updateuser",api.UpdateuserHandler)
	}
}