package initDB

import (
	"auth-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitializeDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/samplego"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	errMigrate := db.AutoMigrate(&models.User{})
	if err != nil {
		panic(errMigrate)
	}
	Database = db
}
