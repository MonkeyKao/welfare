package models

import (
	"time"
	"walfare/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Account  string    `json:"account"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Salt     string    `json:"salt"`
	Birthday time.Time `json:"birthday"`
	Female   int       `json:"female"`
	Location int       `json:"location"`
	Email    string    `json:"email"`
}

func GetUserByID(user *User, id int) error {
	return database.DB.Table("Users").First(user, id).Error
}

func GetUserByAccount(user *User, account string) error {
	return database.DB.Table("Users").Where("account = ?", account).First(user).Error
}

func CreateUser(user *User) error {
	return database.DB.Table("Users").Create(user).Error
}

func UpdateUser(id uint, user *User) error {
	// 使用 GORM 的 Model 方法并通过 map 进行更新
	return database.DB.Table("Users").Where("id = ?", id).Updates(user).Error
}

func DeleteUser(id uint) error {
	return database.DB.Table("Users").Where("id = ?", id).Delete(id).Error
}
