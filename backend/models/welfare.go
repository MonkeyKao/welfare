package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Welfare struct {
	Id              uint     `json:"id"`
	Category        []int    `json:"category"`
	City            int      `json:"city"`
	Date            string   `json:"date"`
	Title           string   `json:"title"`
	Detail          string   `json:"detail"`
	DetailCondition []string `json:"detailCondition"`
	Status          []int    `json:"status"`
	Forward         []string `json:"forward"`
	CanGet          int      `json:"canGet"`
	Url             string   `json:"url"`
}

var welfares []Welfare

func GetWelfares() []Welfare {
	return welfares
}

func GetWelfareByID(id uint) (Welfare, bool) {
	for _, welfare := range welfares {
		if welfare.Id == id {
			return welfare, true
		}
	}
	return Welfare{}, false
}

func GetWelfareFromJson() {
	type data struct {
		Script string    `json:"city"`
		Output []Welfare `json:"output"`
	}

	originalPath, _ := os.Getwd()
	fixedPart := "backend\\cmd"
	index := strings.Index(originalPath, fixedPart)
	newPath := originalPath[:index]
	newPath = fmt.Sprintf("%s\\data.json", newPath)

	fileData, _ := os.ReadFile(newPath)

	jsonStr := string(fileData)

	var dataList []data
	err := json.Unmarshal([]byte(jsonStr), &dataList)
	if err != nil {
		fmt.Println("解析 JSON 错误：", err)
	}

	var result []Welfare
	for _, data := range dataList {
		result = append(result, data.Output...)
	}
	fmt.Print("取得福利json")

	for _, welfare := range result {
		fmt.Println("Loaded welfare:", welfare)
	}

	welfares = result
}
