package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"time"

	"github.com/vipin030/automan/src/models"
	"github.com/vipin030/automan/src/utils/logger"
	util "github.com/vipin030/automan/src/utils"
)

// FindVehicleTypes returns all vehicle type
// ShowAccount godoc
// @Summary Show all vehicle type
// @Accept  json
// @Produce  json
// @Success 200 {object} models.VehicleType
// @Header 200 {string} Token "qwerty"
// @Router /vtypes [get]
func FindVehicleTypes(c *gin.Context) {
	vtypes, err := models.FindAll()
	if err != nil {
		logger.Error("Database error on find vehicle type", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
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
	vtype, err := models.Find(c.Param("id"))
	if err != nil {
		logger.Error("Error on FindVehicleType func ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": vtype})
}

// CreateVehicleType create a new vehicle type
// @Router /vtypes [post]
func CreateVehicleType(c *gin.Context) {
	// Validate input
	input := &models.VehicleType{}
	if err := c.ShouldBindJSON(input); err != nil {
		logger.Error("Error on CreateVehicleType ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicleType := models.VehicleType{
		Name:      input.Name,
		UserID:    input.UserID,
		CreatedAt: util.GetNow(),
	}
	data := vehicleType.Create()

	if status := data["status"].(bool); !status {
		c.JSON(http.StatusBadRequest, gin.H{"error": data["message"]})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Set a recurring reminder to repeat every quarter for postgres minor upgrade
// Create a discovery ticket to investigate how other teams are dealing with migration and what are the benifits we can achieve if we upgrade to latest major version

// UpdateVehicleType Updates specific vehicle type
// @Router /vtypes/{id} [patch]
func UpdateVehicleType(c *gin.Context) {
	// Validate input
	input := &models.VehicleType{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UpdatedAt = util.GetNow()
	_, err := input.Update(c.Param("id"))
	if err != nil {
		logger.Error("Error on UpdateVehicleType ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}

// DeleteVehicleType delete specific vehicle type
// @Router /vtypes/{id} [detete]
func DeleteVehicleType(c *gin.Context) {
	// Get model if exist
	vehicleType := &models.VehicleType{}
	_, err := vehicleType.Delete(c.Param("id"))
	if err != nil {
		logger.Error("Error on DeleteVehicleType ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
