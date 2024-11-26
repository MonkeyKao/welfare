package models

import (
	"walfare/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Account  string   `json:"account"`
	Password string   `json:"-"`
	Name     string   `json:"name"`
	Salt     string   `json:"-"`
	Birthday string   `json:"birthday"`
	Female   int      `json:"female"`
	Location int      `json:"location"`
	Email    string   `json:"email"`
	Families []Family `gorm:"many2many:family_members;"`
}

func (user *User) GetUserByID(id uint) error {
	return database.DB.Preload("Families").First(user, id).Error
}

func (user *User) GetUserByAccount(account string) error {
	return database.DB.Where("account = ?", account).First(user).Error
}

func (user *User) GetUserByEmail(email string) error {
	return database.DB.Where("email = ?", email).First(user).Error
}

func (user *User) GetUserEmailByUserID(id uint) (string, error) {
	var email string
	err := database.DB.Model(&User{}).Select("email").Where("id = ?", id).Take(&email).Error
	if err != nil {
		return "", err
	}
	return email, nil
}

func (user *User) CreateUser() error {
	return database.DB.Create(user).Error
}

func (user *User) UpdateUser() error {
	// 使用 GORM 的 Model 方法并通过 map 进行更新
	return database.DB.Updates(user).Error
}

func (user *User) DeleteUser() error {
	return database.DB.Delete(user).Error
}
