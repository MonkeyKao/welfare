package api

import (
	"fmt"
	"net/http"
	"path/filepath"

	//"fmt"
	"encoding/json"
	"os"
	"walfare/models"

	"github.com/gin-gonic/gin"
)

// 處理JSON數據
func WelfareHandler(c *gin.Context) {
	var welfareList []models.Welfare

	// 取得當前工作目錄
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 使用相對路徑構建完整路徑
	relativePath := filepath.Join(dir, "source.json")
	fmt.Println("Relative Path:", relativePath)

	// 讀取 JSON 文件
	file, err := os.Open(relativePath) // 确保路径正确
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

func GetQAHandler(c *gin.Context) {
	type QA struct {
		Id       uint
		Question string
		Answer   string
	}

	qaList := []QA{
		{Id: 1, Question: "What is the capital of France?", Answer: "Paris"},
		{Id: 2, Question: "What is 2 + 2?", Answer: "4"},
	}

	c.IndentedJSON(http.StatusOK, qaList)
}
