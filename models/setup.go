package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	database, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err !=nil {
		panic("Connection to Database Failed!!")
	}

	err = database.AutoMigrate(&Todo{})

	if err != nil {
		return
	}

	DB = database
}