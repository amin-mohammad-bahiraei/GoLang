package database

import (
	"GoLang/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func init() {

}

func DB() *gorm.DB {
	dsn := "host=localhost user=postgres password=@min1375Amin dbname=GoLang port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Error")
	}
	db.AutoMigrate(&model.Foo{})
	return db
}
