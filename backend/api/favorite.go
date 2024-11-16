package api

import (
	"net/http"
	"strconv"
	"walfare/models"

	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	var welfare []models.Welfare
	userID := c.GetUint("UserID")
	welfare, err := models.GetFavoritesByUserID(userID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, welfare)
}

func AddFavorite(c *gin.Context) {
	welfareIdStr := c.Param("id")
	userID := c.GetUint("UserID")
	// 将 welfareId 从字符串转换为 uint 类型
	welfareId, err := strconv.ParseUint(welfareIdStr, 10, 64)
	if err != nil {
		// 如果转换失败，返回错误
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid welfare ID"})
		return
	}

	if err := models.AddFavorite(uint(welfareId), userID); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"msg": "添加成功"})
}

func DeleteFavorite(c *gin.Context) {
	welfareIdStr := c.Param("id") // 获取 welfare ID
	userID := c.GetUint("UserID") // 获取 UserID

	// 将 welfareId 从字符串转换为 uint 类型
	welfareId, err := strconv.ParseUint(welfareIdStr, 10, 64)
	if err != nil {
		// 如果转换失败，返回错误
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid welfare ID"})
		return
	}

	// 调用模型方法删除收藏
	if err := models.DeleteFavorite(uint(welfareId), userID); err != nil {
		// 如果发生错误，返回 400 错误和错误信息
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回成功消息
	c.IndentedJSON(http.StatusOK, gin.H{"msg": "删除成功"})
}
