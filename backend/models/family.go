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
		FamilyID   uint   `json:"family_id"`
		FamilyName string `json:"family_name"`
		UserName   string `json:"user_name"`
		Role       string `json:"role"`
	}
	if err := database.DB.Table("family_members").
		Select("families.id AS family_id, families.family_name AS family_name, users.name AS user_name, family_members.role").
		Joins("JOIN users ON users.id = family_members.user_id AND users.deleted_at IS NULL").
		Joins("JOIN families ON families.id = family_members.family_id AND families.deleted_at IS NULL").
		Where("family_members.family_id IN ?", familyIDs).
		Find(&rawResults).Error; err != nil {
		return "", err
	}

	// 3. 整理數據為分組結構
	groupedResults := make(map[uint]map[string]interface{}) // key 為 FamilyID
	for _, result := range rawResults {
		// 如果該 FamilyID 尚未初始化，創建一個新的分組
		if _, exists := groupedResults[result.FamilyID]; !exists {
			groupedResults[result.FamilyID] = map[string]interface{}{
				"familyName": result.FamilyName,
				"familyId":   result.FamilyID,
				"users":      []map[string]string{},
			}
		}
		// 添加用戶信息到對應的家庭
		groupedResults[result.FamilyID]["users"] = append(
			groupedResults[result.FamilyID]["users"].([]map[string]string),
			map[string]string{
				"userName": result.UserName,
				"role":     result.Role,
			},
		)
	}

	// 4. 將分組結果轉換為列表
	finalResults := []map[string]interface{}{}
	for _, family := range groupedResults {
		finalResults = append(finalResults, family)
	}

	// 5. 將最終結果轉換為 JSON 字符串
	jsonData, err := json.Marshal(finalResults)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func (family *Family) DeleteFamily() error {
	return database.DB.Delete(family).Error
}
