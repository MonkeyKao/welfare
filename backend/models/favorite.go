package models

import (
	"walfare/database"
)

type Favorite struct {
	ID        uint `gorm:"primaryKey"`
	WelfareID uint `json:"welfare_id"`
	UserID    uint `json:"user_id"`
}

func AddFavorite(welfareID, userID uint) error {
	favorite := Favorite{
		WelfareID: welfareID,
		UserID:    userID,
	}
	return database.DB.Create(&favorite).Error
}

func DeleteFavorite(welfareID, userID uint) error {
	// 嘗試根據 welfare_id 和 user_id 查找並刪除紀錄
	return database.DB.Where("welfare_id = ? AND user_id = ?", welfareID, userID).Delete(&Favorite{}).Error
}

func GetFavoritesByUserID(userID uint) ([]Welfare, error) {
	var favorites []Favorite
	// 查詢所有該 user_id 的 Favorite 記錄
	err := database.DB.Where("user_id = ?", userID).Find(&favorites).Error
	if err != nil {
		return nil, err
	}

	var welfares []Welfare
	allWelfares := GetWelfares() // 獲取所有 Welfare 資料

	// 根據 favorite 的 welfare_id 查詢所有的 Welfare
	for _, favorite := range favorites {
		// 遍歷所有的 Welfare 資料，根據 welfare_id 查找匹配的 Welfare
		for _, welfare := range allWelfares {
			if welfare.Id == int(favorite.WelfareID) { // 如果 Welfare 的 id 與 favorite 的 welfare_id 匹配
				welfares = append(welfares, welfare) // 加入到 welfares 切片
				break                                // 找到後就停止尋找，避免重複加入
			}
		}
	}

	return welfares, nil
}
