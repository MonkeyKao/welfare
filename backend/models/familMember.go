package models

import "walfare/database"

type FamilyMember struct {
	UserID   uint `gorm:"primaryKey"`
	FamilyID uint `gorm:"primaryKey"`
	Role     uint `gorm:"size:50"`
}

func (familyMember *FamilyMember) CreateFamilyMember() error {
	return database.DB.Create(familyMember).Error
}

func (familyMember *FamilyMember) DeleteFamilyMember() error {
	return database.DB.Delete(familyMember).Error
}
