package database

import (
	"github.com/ahmed-deftoner/ambassador/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "admin:Ghtwhts786@tcp(database-1.ckpee9qqxzv9.ap-south-1.rds.amazonaws.com:3306)/db"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to mysql")
	}
}

func AddMigration() {
	DB.AutoMigrate(models.User{})
}
