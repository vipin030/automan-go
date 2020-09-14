package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	//dsn := "host=db port=5433 user=interface password=interface dbname=employee sslmode=disable"
	dsn := "postgres://interface:interface@db/employee?sslmode=disable"
	fmt.Println(dsn)
	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{}, &VehicleType{})

	db.LogMode(true)

	DB = db
}
