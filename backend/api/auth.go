package api

import (
	"fmt"
	"net/http"

	"walfare/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.Request.Header.Get("token")

		if authHandler == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未攜帶token"})
			c.Abort()
			return
		}

		//调用下方自己实现的token解析函数,并且在判断token是否过期
		mc, err := utils.ParseToken(authHandler)
		if err != nil {
			fmt.Println("err = ", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"isLogin": false,
				"msg":     "无效的Token",
			})
			c.Abort()
			return
		}

		// 确保将 mc.UserID 转换为 int 类型
		userID := uint(mc.UserID) // 如果 mc.UserID 是 int64 或 uint，进行类型转换
		// 将当前请求的 userID 信息保存到请求的上下文 c 上
		c.Set("UserID", userID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}

}
