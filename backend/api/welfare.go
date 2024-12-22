package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"walfare/models"
	"walfare/services"

	//"fmt"

	"github.com/gin-gonic/gin"
)

func GetWelfareHandler(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))

	// 获取查询参数
	cityStrings := c.QueryArray("city[]") // 支持 city[] 格式
	title := c.Query("title")
	categoryStrings := c.QueryArray("category[]")
	statusStrings := c.QueryArray("status[]")

	// 将字符串数组转换为整数数组
	var cityFilter []int
	var categoryFilter []int
	var statusFilter []int
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
	for _, cs := range statusStrings {
		if id, err := strconv.Atoi(cs); err == nil {
			statusFilter = append(statusFilter, id)
		}
	}

	welfares := services.Filiter(cityFilter, categoryFilter, statusFilter, title)

	// 分页处理
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(welfares) {
		welfares = []services.WelfareResponse{}
	} else if end > len(welfares) {
		welfares = welfares[start:]
	} else {
		welfares = welfares[start:end]
	}

	c.IndentedJSON(http.StatusOK, welfares)
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
