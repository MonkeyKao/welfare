package api

import (
	"net/http"
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

	qaList := []QA{
		{Id: 1, Question: "What is the capital of France?", Answer: "Paris"},
		{Id: 2, Question: "What is 2 + 2?", Answer: "4"},
	}

	c.IndentedJSON(http.StatusOK, qaList)
}
