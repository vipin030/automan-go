package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres driver
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

// DB object
var DB *gorm.DB

func init() {
	dir, err := filepath.Abs(".env")
	if err != nil {
		log.Fatal(err)
	}
	if err := godotenv.Load(dir); err != nil {
		log.Fatal("Error loading .env file", err)
	}
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
