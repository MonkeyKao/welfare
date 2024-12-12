package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"walfare/models"

	//"fmt"

	"github.com/gin-gonic/gin"
)

// 處理JSON數據
func WelfareHandler(c *gin.Context) {
	type WelfareResponse struct {
		Id       uint   `json:"id"`
		Title    string `json:"title"`
		City     string `json:"city"`
		Category []int  `json:"category"`
		CanGet	 int	`json:"canGet"`
	}

	var response []WelfareResponse
	welfares := models.GetWelfares()
	for _, welfare := range welfares {
		response = append(response, WelfareResponse{
			Id:       welfare.Id,
			City:     welfare.City,
			Category: welfare.Category,
			Title:    welfare.Title,
			CanGet:   welfare.CanGet,
		})
	}
	c.IndentedJSON(http.StatusOK, response)
}

func GetWelfareByID(c *gin.Context) {
	welfareID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if welfare, exitis := models.GetWelfareByID(uint(welfareID)); exitis {
		c.IndentedJSON(http.StatusOK, welfare)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
	}
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
