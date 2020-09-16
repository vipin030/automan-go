package models

// User model
type User struct {
	ID       uint64 `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
