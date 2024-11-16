package models

type FamilyMember struct {
	UserID   uint   `gorm:"primaryKey"`
	FamilyID uint   `gorm:"primaryKey"`
	Role     string `gorm:"size:50"`
}
