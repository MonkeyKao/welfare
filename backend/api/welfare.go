package api

import (
	"net/http"

	//"fmt"
	"encoding/json"
	"os"
	"walfare/models"

	"github.com/gin-gonic/gin"
)

// 處理JSON數據
func WelfareHandler(c *gin.Context) {
	var welfareList []models.Welfare

	// 讀取 JSON 文件
	file, err := os.Open("C:\\Users\\User\\Desktop\\welfare\\backend\\api\\source.json") // 确保路径正确
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// 解析 JSON
	if err := json.NewDecoder(file).Decode(&welfareList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回 JSON 數據
	c.JSON(http.StatusOK, welfareList)
}
