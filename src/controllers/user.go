package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/vipin030/automan/src/models"
)

// CreateUserAccount creates new user
func CreateUserAccount(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON Provided")
		return
	}
	resp, err := user.Create()

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

// Authenticate using user credentials
func Authenticate(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided")
		return
	}
	resp, error := models.Login(user.Email, user.Password)
	if error != nil {
		c.JSON(error.Status(), gin.H{"error": error.Message()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": resp["token"], "user": resp["user"]})
}
