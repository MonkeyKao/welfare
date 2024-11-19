package models

import (
	"encoding/json"
	"walfare/database"

	"gorm.io/gorm"
)

type Family struct {
	gorm.Model
	FamilyName string `gorm:"not null"`

	// 關聯: 一個家庭可以有多個用戶
	Members []User `gorm:"many2many:family_members;"`
}

func (family *Family) CreateFamily() error {
	return database.DB.Create(family).Error
}

// 大便代碼！！！請勿觸動
func GetAllFamilyByUserID(userID uint) (string, error) {
	// 1. 找出用戶所屬的所有家庭 ID
	var familyIDs []uint
	if err := database.DB.Table("family_members").
		Select("family_id").
		Where("user_id = ?", userID).
		Find(&familyIDs).Error; err != nil {
		return "", err
	}

	// 2. 查詢這些家庭中的所有成員及其角色
	var rawResults []struct {
		FamilyName string `json:"family_name"`
		UserName   string `json:"user_name"`
		Role       string `json:"role"`
	}
	if err := database.DB.Table("family_members").
		Select("families.family_name AS family_name, users.name AS user_name, family_members.role").
		Joins("JOIN users ON users.id = family_members.user_id").
		Joins("JOIN families ON families.id = family_members.family_id").
		Where("family_members.family_id IN ?", familyIDs).
		Find(&rawResults).Error; err != nil {
		return "", err
	}

	// 3. 整理數據為分組結構
	groupedResults := make(map[string][]map[string]string)
	for _, result := range rawResults {
		groupedResults[result.FamilyName] = append(groupedResults[result.FamilyName], map[string]string{
			"user_name": result.UserName,
			"role":      result.Role,
		})
	}

	// 4. 將分組結果轉換為列表
	finalResults := []map[string]interface{}{}
	for familyName, users := range groupedResults {
		finalResults = append(finalResults, map[string]interface{}{
			"family_name": familyName,
			"users":       users,
		})
	}

	// 5. 將最終結果轉換為 JSON 字符串
	jsonData, err := json.Marshal(finalResults)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
