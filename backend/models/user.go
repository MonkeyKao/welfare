package models

import (
	"time"
	"walfare/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Account  string    `json:"account"`
	Password string    `json:"-"`
	Name     string    `json:"name"`
	Salt     string    `json:"-"`
	Birthday time.Time `json:"birthday"`
	Female   int       `json:"female"`
	Location int       `json:"location"`
	Email    string    `json:"email"`
}

func (user *User) GetUserByID(id uint) error {
	return database.DB.Table("Users").First(user, id).Error
}

func (user *User) GetUserByAccount(account string) error {
	return database.DB.Table("Users").Where("account = ?", account).First(user).Error
}

func (user *User) GetUserByEmail(email string) error {
	return database.DB.Table("Users").Where("email = ?", email).First(user).Error
}

func (user *User) CreateUser() error {
	return database.DB.Table("Users").Create(user).Error
}

func (user *User) UpdateUser() error {
	// 使用 GORM 的 Model 方法并通过 map 进行更新
	return database.DB.Table("Users").Where("id = ?", user.ID).Updates(user).Error
}

func (user *User) DeleteUser() error {
	return database.DB.Table("Users").Where("id = ?", user.ID).Delete(user.ID).Error
}
