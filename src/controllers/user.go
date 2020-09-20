package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/vipin030/automan/src/models"
)

// Authenticate using user credentials
func Authenticate(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided")
		return
	}
	resp, token, err := models.Login(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "user": resp})
}
