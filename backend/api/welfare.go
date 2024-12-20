package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"walfare/models"

	//"fmt"

	"github.com/gin-gonic/gin"
)

// 處理JSON數據
func GetWelfareHandler(c *gin.Context) {
	type WelfareResponse struct {
		Id       uint   `json:"id"`
		Title    string `json:"title"`
		City     string `json:"city"`
		Category []int  `json:"category"`
		CanGet   int    `json:"canGet"`
	}

	var response []WelfareResponse

	// 获取查询参数
	city := c.Query("city")
	title := c.Query("title")
	categoryStr := c.QueryArray("category")

	// 转换 category 参数为整数切片
	var categoryFilter []int
	for _, c := range categoryStr {
		categoryFilter = append(categoryFilter, atoi(c))
	}

	welfares := models.GetWelfares()
	// 结果过滤
	for _, welfare := range welfares {
		// 检查筛选条件
		if (city == "" || strings.EqualFold(welfare.City, city)) &&
			(title == "" || strings.Contains(strings.ToLower(welfare.Title), strings.ToLower(title))) &&
			(len(categoryFilter) == 0 || hasCategory(welfare.Category, categoryFilter)) {
			response = append(response, WelfareResponse{
				Id:       welfare.Id,
				Title:    welfare.Title,
				City:     welfare.City,
				Category: welfare.Category,
				CanGet:   0,
			})
		}
	}

	c.IndentedJSON(http.StatusOK, response)
}

func hasCategory(welfareCategories, filterCategories []int) bool {
	categoryMap := make(map[int]bool)
	for _, c := range filterCategories {
		categoryMap[c] = true
	}
	for _, wc := range welfareCategories {
		if categoryMap[wc] {
			return true
		}
	}
	return false
}

// 安全地将字符串转换为整数
func atoi(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
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
