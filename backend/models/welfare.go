package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Script string `json:"script"`
	Output string `json:"output"`
}
type Welfare struct {
	Id              int    `json:"id"`
	Category        []int  `json:"category"`
	City            string `json:"city"`
	Date            string `json:"date"`
	Priority        int    `json:"priority"`
	Title           string `json:"title"`
	Detail          string `json:"detail"`
	DetailCondition string `json:"detailCodition"`
	DetailDocument  string `json:"detailDocument"`
	DetailLink      string `json:"url"`
}

func ParseData(jsonStr string) []Welfare {
	var dataList []Data
	err := json.Unmarshal([]byte(jsonStr), &dataList)
	if err != nil {
		fmt.Println("解析 JSON 错误：", err)
		return nil
	}

	var result []Welfare
	for _, data := range dataList {
		var welfare []Welfare
		err := json.Unmarshal([]byte(data.Output), &welfare)
		if err != nil {
			fmt.Println("解析 Output 错误：", err)
			continue
		}
		result = append(result, welfare...)
	}
	return result
}

func GetWelfares() []Welfare {
	fmt.Println(os.Getwd())
	fileData, err := os.ReadFile("C:\\flutter\\welfare\\backend\\cmd\\data.json")

	if err != nil {
		fmt.Println("Error reading file:", err)

	}

	jsonStr := string(fileData)

	items := ParseData(jsonStr)
	return items
}
