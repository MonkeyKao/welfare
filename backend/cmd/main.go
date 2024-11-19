package main

import (
	"walfare/api"
	"walfare/database"
	"walfare/routes"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "user=root password=123456 host=43.163.2.11 port=5432 dbname=walfare sslmode=disable" //資料庫
	database.InitDB(dsn)                                                                         //去連接資料庫

	router := gin.Default() //初始化路由
	routes.UserRoutes(router)
	routes.WelfareRoutes(router)
	routes.VerifyRoutes(router)
	routes.FavoriteRoutes(router)
	routes.FamilyRoutes(router)
	router.GET("QA", api.GetQAHandler)
	router.Run(":8081")
}
