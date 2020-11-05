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
	"github.com/vipin030/automan/src/utils/errors"
)

// User model
type User struct {
	ID       uint64 `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	Phone    string `json:"phone"`
}

// Validate the new user
func (user *User) Validate() errors.APIError {
	if !strings.Contains(user.Email, "@") {
		return errors.ValidationError("Email address is required", "validation_error")
	}
	if len(user.Password) < 6 {
		return errors.ValidationError("Password is not Strong", "password_not_strong")
	}

	temp := &User{}
	err := DB.Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Error: ", err)
		return errors.InternalServerError("Connection failed", err)
	}
	if temp.Email != "" {
		return errors.ValidationError("Email already exist", "email_record_exist")
	}
	return nil
}

// Create a new user
func (user *User) Create() (map[string]interface{}, errors.APIError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	password, err := GeneratePass(user.Password)
	if err != nil {
		return nil, errors.InternalServerError("Password Hash generation failed", err)
	}
	user.Password = password

	if err := DB.Create(user).Error; err != nil {
		errText := "Failed to create a user"
		log.Println(errText, err)
		return nil, errors.InternalServerError(errText, err)
	}
	return util.Message(true, "Account has been created"), nil
}

// GeneratePass return hash generated password
var GeneratePass = func(password string) (string, error) {
	passwordHash, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), error
}

// Login function
func Login(email, password string) (map[string]interface{}, errors.APIError) {

	user := &User{}
	if err := DB.Debug().Select("id, email, password, phone").Where("email = ?", email).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFoundError("Email not found")
		}
		return nil, errors.InternalServerError("Connection error", err)
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return nil, errors.UnAutherizedError("Invalid login credentials. Please try again")
	}
	user.Password = ""
	token, err := CreateToken(user.ID)
	if err != nil {
		return nil, errors.InternalServerError("Token Creation Failed", err)
	}
	resp := util.Message(true, "Success")
	resp["token"] = token
	resp["user"] = user
	return resp, nil
}

// CreateToken create new token
func CreateToken(UserID uint64) (string, errors.APIError) {
	var err error
	os.Setenv("ACCESS_SECRET", "lkngdogop")

	atClaims := jwt.MapClaims{}
	atClaims["autherized"] = true
	atClaims["user_id"] = UserID
	atClaims["exp"] = util.GetNow().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", errors.InternalServerError("Token Creation Failed", err)
	}
	return token, nil
}
