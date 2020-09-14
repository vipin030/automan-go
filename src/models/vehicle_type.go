package models

import (
	"time"
)

type VehicleType struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	UserId    uint64    `gorm:"TYPE:integer REFERENCES users"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
