package database

import (
	"first-application/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("user=postgres password=root dbname=golang-test port=5432"), &gorm.Config{})

	if err != nil {
		panic("Could not cannot to the database")
	}

	fmt.Println(db)
	DB = db
	db.AutoMigrate(&models.User{})
}
