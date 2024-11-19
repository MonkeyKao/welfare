package models

import "walfare/database"

type FamilyMember struct {
	UserID   uint   `gorm:"primaryKey"`
	FamilyID uint   `gorm:"primaryKey"`
	Role     string `gorm:"size:50"`
}

func (familyMember *FamilyMember) CreateFamilyMember() error {
	return database.DB.Create(familyMember).Error
}
