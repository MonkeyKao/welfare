package database

import(
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string){
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) //因為DB,err的類型是知道的，所以先訂下類型用=
	if err != nil{ 
		log.Fatal("連接失敗! , 原因 : %v", err)
	}
}