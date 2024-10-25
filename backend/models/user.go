package models

import(
	"gorm.io/gorm"
	"walfare/database"
	"time"
)

type User struct{
	gorm.Model
	
	Account 	string 		`json:"account"`
	Password 	string 		`json:"password"`
	Name		string		`json:"name"`
	Salt 		string 		`json:"salt"`
	Birthday 	time.Time 	`json:"birthday"`
	Female 		int 		`json:"female"`
	Location 	int 		`json:"location"`
	Email 		string 		`json:"email"`
}

func GetUserByID(user *User, id int)error{
	return database.DB.Table("Users").First(user,id).Error
}

func GetUserByAccount(user *User,account string)error{
	return database.DB.Table("Users").Where("account = ?",account).First(user).Error
}

func CreateUser(user *User)error {
	return database.DB.Table("Users").Create(user).Error
}

