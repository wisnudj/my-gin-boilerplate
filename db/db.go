package db

import (
	"my-gin-boilerplate/config"
	"my-gin-boilerplate/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func Init() {
	config := config.GetConfig()

	dialect := config.GetString("db.dialect")
	url := config.GetString("db.url")

	db, _ = gorm.Open(dialect, url)

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
