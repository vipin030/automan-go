package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres driver
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// DB object
var DB *gorm.DB

func init() {
	dir, err := filepath.Abs(".env")
	if err != nil {
		log.Fatal(err)
	}
	dir = strings.Replace(dir, "/models", "", 1)
	if err := godotenv.Load(dir); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	var (
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		host     = os.Getenv("POSTGRES_HOST")
		schema   = os.Getenv("POSTGRES_DB")
	)
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", username, password, host, schema)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&User{}, &VehicleType{})
	db.LogMode(true)
	DB = db
}
