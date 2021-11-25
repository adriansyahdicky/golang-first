package database

import (
	"first-application/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	db, err := gorm.Open(postgres.Open("user=postgres password=root dbname=golang-test port=5432"), &gorm.Config{})

	if err != nil {
		panic("Could not cannot to the database")
	}

	fmt.Println(db)
	db.AutoMigrate(&models.User{})
}
