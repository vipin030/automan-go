package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vipin030/automan/models"
)

type VehicleType struct {
	Name        string    `json:"name" binding:"required"`
	UserId      uint64    `json:"user_id" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateVehicleTypeInput struct {
	Name        string    `json:"name"`
	UserId      uint64    `json:"user_id" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FindVehicleTypes(c *gin.Context) {
	var vtypes []models.VehicleType
	models.DB.Find(&vtypes)

	c.JSON(http.StatusOK, gin.H{"data": vtypes})
}

func FindVehicleType(c *gin.Context) {
	// Get model if exist
	var vtype models.VehicleType
	if err := models.DB.Where("id = ?", c.Param("id")).First(&vtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vtype})
}

func CreateVehicleType(c *gin.Context) {
	// Validate input
	var input VehicleType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle_type := models.VehicleType{Name: input.Name, UserId: input.UserId, CreatedAt: time.Now().UTC()}
	if err := models.DB.Create(&vehicle_type); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vehicle_type})
}

func UpdateVehicleType(c *gin.Context) {
	// Get model if exist
	var vehicle_type models.VehicleType
	if err := models.DB.Where("id = ?", c.Param("id")).First(&vehicle_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateVehicleTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UpdatedAt = time.Now().UTC()

	models.DB.Model(&vehicle_type).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": vehicle_type})
}

func DeleteVehicleType(c *gin.Context) {
	// Get model if exist
	var vehicle_type models.VehicleType
	if err := models.DB.Where("id = ?", c.Param("id")).First(&vehicle_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&vehicle_type)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
