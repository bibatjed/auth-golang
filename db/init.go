package initDB

import (
	"auth-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Database Dbinstance

func InitializeDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/samplego"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	Database.Db = db
}
