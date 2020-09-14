package models

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	//dsn := "postgres://user:password@db/db?sslmode=disable"
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DB"));
	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{}, &VehicleType{})

	db.LogMode(true)

	DB = db
}
