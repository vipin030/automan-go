package controllers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"

	"github.com/vipin030/automan/src/models"
)

// User struct
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone"`
}

// UserList struct
type UserList struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"_"`
	Phone    string `json:"phone"`
}

// Authenticate using user credentials
func Authenticate(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided")
		return
	}

	var user User
	if err := models.DB.Debug().Select("id, username, phone").Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide valid login details"})
		return
	}

	token, err := CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
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
