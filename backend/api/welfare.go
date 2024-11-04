package api

import (
	"encoding/json"
	"net/http"
	"os"
	"walfare/models"

	//"fmt"

	"github.com/gin-gonic/gin"
)

// 處理JSON數據
func WelfareHandler(c *gin.Context) {
	items := models.GetWelfares()
	c.IndentedJSON(http.StatusOK, items)
}

func GetQAHandler(c *gin.Context) {
	type QA struct {
		Id       uint
		Question string
		Answer   string
	}

	fileData, _ := os.ReadFile("../../QA.json")
	jsonStr := string(fileData)

	var qaList []QA
	err := json.Unmarshal([]byte(jsonStr), &qaList)

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, qaList)
}
