package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/vipin030/automan/src/models"
)

// VehicleType struct
type VehicleType struct {
	Name      string    `json:"name" binding:"required"`
	UserID    uint64    `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateVehicleTypeInput struct
type UpdateVehicleTypeInput struct {
	Name      string    `json:"name"`
	UserID    uint64    `json:"user_id" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FindVehicleTypes returns all vehicle type
// ShowAccount godoc
// @Summary Show all vehicle type
// @Accept  json
// @Produce  json
// @Success 200 {object} models.VehicleType
// @Header 200 {string} Token "qwerty"
// @Router /vtypes [get]
func FindVehicleTypes(c *gin.Context) {
	var vtypes []models.VehicleType
	models.DB.Find(&vtypes)

	c.JSON(http.StatusOK, gin.H{"data": vtypes})
}

// FindVehicleType return single vehicle type
// ShowAccount godoc
// @Summary Show a vehicle type
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} models.VehicleType
// @Header 200 {string} Token "qwerty"
// @Router /vtypes/{id} [get]
func FindVehicleType(c *gin.Context) {
	// Get model if exist
	var vtype models.VehicleType
	if err := models.DB.Where("id = ?", c.Param("id")).First(&vtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vtype})
}

// CreateVehicleType create a new vehicle type
// @Router /vtypes [post]
func CreateVehicleType(c *gin.Context) {
	// Validate input
	var input VehicleType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicleType := models.VehicleType{
		Name:      input.Name,
		UserID:    input.UserID,
		CreatedAt: time.Now().UTC(),
	}
	if err := models.DB.Create(&vehicleType); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vehicleType})
}

// UpdateVehicleType Updates specific vehicle type
// @Router /vtypes/{id} [patch]
func UpdateVehicleType(c *gin.Context) {
	// Get model if exist
	var vehicleType models.VehicleType
	if err := models.DB.Where("id = ?", c.Param("id")).First(&vehicleType).Error; err != nil {
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

	models.DB.Model(&vehicleType).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": vehicleType})
}

// DeleteVehicleType delete specific vehicle type
// @Router /vtypes/{id} [detete]
func DeleteVehicleType(c *gin.Context) {
	// Get model if exist
	var vehicleType models.VehicleType
	if err := models.DB.Where("id = ?", c.Param("id")).First(&vehicleType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&vehicleType)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
