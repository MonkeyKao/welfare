package services

import (
	"strings"
	"walfare/models"
)

type WelfareResponse struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	City     int    `json:"city"`
	Category []int  `json:"category"`
	CanGet   int    `json:"canGet"`
}

// 篩選器
func Filiter(location []int, service []int, status []int, title string) []WelfareResponse {

	welfares := models.GetWelfares()

	var filiterWelfare []WelfareResponse
	for _, welfare := range welfares {
		// 篩選地區 && 篩選服務 && 篩選標題
		if (len(location) == 0 || intInSlice(welfare.City, location)) &&
			(len(service) == 0 || hasCategory(welfare.Category, service)) &&
			(title == "" || strings.Contains(welfare.Title, title)) {
			CanGet := 3

			for _, item := range welfare.Status {
				for _, statu := range status {
					if statu == item {
						CanGet = 1
					}
				}

				if item == 12 {
					CanGet = 2
				}
			}

			filiterWelfare = append(filiterWelfare, WelfareResponse{
				Id:       welfare.Id,
				Title:    welfare.Title,
				City:     welfare.City,
				Category: welfare.Category,
				CanGet:   CanGet,
			})
		}
	}
	return filiterWelfare
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
