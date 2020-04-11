package db

import (
	"my-gin-boilerplate/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func Init() {
	db, _ = gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.People{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
