package config

import (
	"GoLang/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init() {
	var err error
	Database, err = gorm.Open(postgres.Open("host=localhost user=postgres password=@min1375Amin dbname=GoLang port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error Database: %s \n", err)
	} else {
		Database.AutoMigrate(&models.User{})
		Database.AutoMigrate(&models.Foo{})
	}
	fmt.Println("Database connection successful.")
}
