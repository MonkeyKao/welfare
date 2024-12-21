package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"walfare/models"

	//"fmt"

	"github.com/gin-gonic/gin"
)

func GetWelfareHandler(c *gin.Context) {
	type WelfareResponse struct {
		Id       uint   `json:"id"`
		Title    string `json:"title"`
		City     int    `json:"city"`
		Category []int  `json:"category"`
		CanGet   int    `json:"canGet"`
	}

	var response []WelfareResponse
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))

	// 获取查询参数
	cityStrings := c.QueryArray("city[]") // 支持 city[] 格式
	title := c.Query("title")
	categoryStrings := c.QueryArray("category[]")

	// 将字符串数组转换为整数数组
	var cityFilter []int
	var categoryFilter []int
	for _, cs := range cityStrings {
		if id, err := strconv.Atoi(cs); err == nil {
			cityFilter = append(cityFilter, id)
		}
	}
	for _, cs := range categoryStrings {
		if id, err := strconv.Atoi(cs); err == nil {
			categoryFilter = append(categoryFilter, id)
		}
	}

	welfares := models.GetWelfares()
	// 过滤数据
	filtered := []WelfareResponse{}
	for _, welfare := range welfares {
		// 筛选逻辑：城市需要匹配任意一个 cityFilter 中的值
		if (len(cityFilter) == 0 || intInSlice(welfare.City, cityFilter)) &&
			(title == "" || strings.Contains(strings.ToLower(welfare.Title), strings.ToLower(title))) &&
			(len(categoryFilter) == 0 || hasCategory(welfare.Category, categoryFilter)) {
			filtered = append(filtered, WelfareResponse{
				Id:       welfare.Id,
				Title:    welfare.Title,
				City:     welfare.City,
				Category: welfare.Category,
				CanGet:   0,
			})
		}
	}

	// 分页处理
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(filtered) {
		response = []WelfareResponse{}
	} else if end > len(filtered) {
		response = filtered[start:]
	} else {
		response = filtered[start:end]
	}

	c.IndentedJSON(http.StatusOK, response)
}

func intInSlice(value int, list []int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
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
