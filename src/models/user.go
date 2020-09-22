package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// User model
type User struct {
	ID       uint64 `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// Login function
func Login(userName, password string) (*User, string, error) {

	user := &User{}
	if err := DB.Debug().Select("id, username, phone").Where("username = ? AND password = ?", userName, password).First(user).Error; err != nil {
		return nil, "", err
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}


// CreateToken create new token
func CreateToken(UserID uint64) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "lkngdogop")

	atClaims := jwt.MapClaims{}
	atClaims["autherized"] = true
	atClaims["user_id"] = UserID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
