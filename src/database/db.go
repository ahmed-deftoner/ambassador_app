package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	dsn := "admin:Ghtwhts786@tcp(database-1.ckpee9qqxzv9.ap-south-1.rds.amazonaws.com:3306)/db"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to mysql")
	}
}
