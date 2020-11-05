package models

import (
	util "github.com/vipin030/automan/src/utils"
	"log"
	"time"

	"github.com/vipin030/automan/src/utils/errors"
)

// VehicleType Model
type VehicleType struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	UserID    uint64    `json:"user_id" gorm:"TYPE:integer REFERENCES users"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate validate the user Input
func (vehicleType *VehicleType) Validate() errors.APIError {

	if vehicleType.Name == "" {
		return errors.ValidationError("Vehicle type should not be empty", "Validation Error")
	}
	//All the required parameters are present
	return nil
}

// Create return new vehicle type
func (vehicleType *VehicleType) Create() (map[string]interface{}, errors.APIError) {

	error := vehicleType.Validate()
	if error != nil {
		return nil, error
	}
	if err := DB.Create(vehicleType).Error; err != nil {
		log.Println("Error: ", err)
		return nil, errors.InternalServerError("Vehicle Type creation failed", err)
	}
	resp := util.Message(true, "Success")
	resp["data"] = vehicleType
	return resp, nil
}

// Update return upated vehicle type
func (vehicleType *VehicleType) Update(id string) errors.APIError {
	existingVehicleType, err := Find(id)
	if err != nil {
		return errors.BadRequestError("Failure during fetching records")
	}

	if err := DB.Model(&existingVehicleType).Updates(vehicleType).Error; err != nil {
		return errors.InternalServerError("Failure during updating records", err)
	}
	return nil
}

// Delete delete a vehicle type
func (vehicleType *VehicleType) Delete(id string) errors.APIError {

	existingVehicleType, err := Find(id)

	if err != nil {
		return errors.BadRequestError("Failure during fetching")
	}
	if err := DB.Delete(existingVehicleType).Error; err != nil {
		return errors.InternalServerError("Failure duing deleting", err)
	}
	return nil
}

// FindAll returns all vehicle type
func FindAll() ([]VehicleType, errors.APIError) {
	var vtypes []VehicleType
	if err := DB.Find(&vtypes).Error; err != nil {
		return nil, errors.BadRequestError("Failure during fetching")
	}
	return vtypes, nil
}

// Find returns single vehicle type
func Find(id string) (*VehicleType, errors.APIError) {
	vtype := &VehicleType{}
	if err := DB.Where("id = ?", id).First(vtype).Error; err != nil {
		return nil, errors.BadRequestError("Record not found")
	}
	return vtype, nil
}
