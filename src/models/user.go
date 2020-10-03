package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
	"time"

	util "github.com/vipin030/automan/src/utils"
)

// User model
type User struct {
	ID       uint64 `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	Phone    string `json:"phone"`
}

// Validate the new user
func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return util.Message(false, "Email address is required"), false
	}
	if len(user.Password) < 6 {
		return util.Message(false, "Password is not Strong"), false
	}

	temp := &User{}
	err := DB.Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Error: ", err)
		return util.Message(false, "Connection failed"), false
	}
	if temp.Email != "" {
		return util.Message(false, "Email already exist"), false
	}
	return util.Message(false, "Validation successful"), true
}

// Create a new user
func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}
	password, err := GeneratePass(user.Password)
	if err != nil {
		return util.Message(false, "Password Hash generation failed")
	}
	user.Password = password

	if err := DB.Create(user).Error; err != nil {
		log.Println("Error: ", err)
		return util.Message(false, "Failed to create a user")
	}
	return util.Message(true, "Account has been created")
}

// GeneratePass return hash generated password
var GeneratePass = func(password string) (string, error) {
	passwordHash, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), error
}

// Login function
func Login(email, password string) map[string]interface{} {

	user := &User{}
	if err := DB.Debug().Select("id, email, password, phone").Where("email = ?", email).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return util.Message(false, "Email not found")
		}
		log.Println("Error: ", err)
		return util.Message(false, "Connection error")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return util.Message(false, "Invalid login credentials. Please try again")
	}
	user.Password = ""
	token, err := CreateToken(user.ID)
	if err != nil {
		return util.Message(false, "Token creation failed")
	}
	resp := util.Message(true, "Success")
	resp["token"] = token
	resp["user"] = user
	return resp
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
