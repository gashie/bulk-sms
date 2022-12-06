package database

import (
	"bulk-sms/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("mygo:mygo@tcp(192.168.64.2:3306)/go_sms?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database
	// database.Migrator().DropTable(&models.Group{})
	database.AutoMigrate(&models.Group{}, &models.Contact{})
}
