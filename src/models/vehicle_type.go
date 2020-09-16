package models

import (
	"time"
)
// VehicleType Model
type VehicleType struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	UserID    uint64    `gorm:"TYPE:integer REFERENCES users"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
