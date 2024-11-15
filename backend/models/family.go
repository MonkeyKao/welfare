package models

import (
	"time"

	"gorm.io/gorm"
)

type Family struct {
	gorm.Model
	FamilyName string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// 關聯: 一個家庭可以有多個用戶
	Members []User `gorm:"many2many:family_members;"`
}
