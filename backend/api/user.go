package api

import (
	"net/http"
	//"fmt"
	"walfare/models"
	"walfare/services"

	"github.com/gin-gonic/gin"
)

// 登入ing
func DoLoginHandler(c *gin.Context) {
	from := models.User{}

	if err := c.ShouldBind(&from); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	//調用函式去進行驗證，正確會返回一個token
	token, err := services.Login(from.Account, from.Password)

	if err != nil {
		c.IndentedJSON(http.StatusOK, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tokenName": "token", "tokenValue": token})
}

// 註冊
func RegisterHandler(c *gin.Context) {
	form := models.User{}

	if err := c.ShouldBind(&form); err != nil {
		c.IndentedJSON(http.StatusOK, err.Error())
		return
	}

	if err := services.Register(&form); err != nil {
		c.IndentedJSON(http.StatusOK, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, form.ID)
}

// UpdateuserHandler 处理用户更新请求
func UpdateuserHandler(c *gin.Context) {
	// 从上下文中获取用户 ID
	userID := c.GetUint("UserID")

	// 定义一个 map 用于接收更新数据
	var updateData map[string]interface{}

	// 从请求中绑定数据到 map 中
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新用户信息
	if err := models.UpdateUser(userID, updateData); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// 返回成功响应
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
