package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("shortener.db"), &gorm.Config{})
	if err != nil {
		log.Panicln("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return db
}
