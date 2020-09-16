package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres driver
	"os"
)

// DB object
var DB *gorm.DB

// ConnectDatabase create db connection
func ConnectDatabase() {
	//dsn := "postgres://user:password@db/db?sslmode=disable"
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{}, &VehicleType{})

	db.LogMode(true)

	DB = db
}
