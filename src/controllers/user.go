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
	resp := user.Create()
	if status := resp["status"].(bool); !status {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": resp["message"]})
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
	resp := models.Login(user.Email, user.Password)
	if status := resp["status"].(bool); !status {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": resp["message"]})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": resp["token"], "user": resp["user"]})
}
