package api

import (
	"fmt"
	"net/http"
	"time"

	//"fmt"
	"walfare/models"
	"walfare/services"
	"walfare/utils"

	"github.com/gin-gonic/gin"
)

// 登入ing
func DoLoginHandler(c *gin.Context) {
	from := models.User{}

	if err := c.ShouldBindBodyWithJSON(&from); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//調用函式去進行驗證，正確會返回一個token
	token, err := services.Login(from.Account, from.Password)

	// token無效
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tokenName": "token", "tokenValue": token})
}

// 註冊
func RegisterHandler(c *gin.Context) {
	form := models.User{}

	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := services.Register(&form)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"msg": token})
}

func VerifyEmailHandler(c *gin.Context) {
	token := c.Request.Header.Get("token")
	code := c.Param("id")
	fmt.Printf("token:%s", token)
	fmt.Println(code)

	userId, err := services.VerifyEmail(code, token)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const TokenExpireDuration = time.Hour * 24 * 30 * 12 * 30
	newToken, _ := utils.GenerateToken(userId, "", TokenExpireDuration)

	c.IndentedJSON(http.StatusOK, gin.H{"msg": newToken})
}

// UpdateuserHandler 处理用户更新请求
func UpdateuserHandler(c *gin.Context) {
	// 从上下文中获取用户 ID
	userID := c.GetUint("UserID")

	// 定义一个 map 用于接收更新数据
	var form models.User

	// 从请求中绑定数据到 map 中
	if err := c.ShouldBindJSON(&form); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form.ID = userID

	// 更新用户信息
	if err := models.UpdateUser(userID, &form); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// 返回成功响应
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func GetUserByUserIDHandler(c *gin.Context) {
	// 从上下文中获取用户 ID
	userID := c.GetUint("UserID")

	var user models.User
	if err := models.GetUserByID(&user, int(userID)); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
