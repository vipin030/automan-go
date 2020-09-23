package models

import (
	"time"
	util "github.com/vipin030/automan/src/utils"
)
// VehicleType Model
type VehicleType struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	UserID    uint64    `gorm:"TYPE:integer REFERENCES users"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate validate the user Input
func (vehicleType *VehicleType) Validate() (map[string]interface{}, bool) {

	if vehicleType.Name == "" {
		return util.Message(false, "Vehicle type should not be empty"), false
	}
	//All the required parameters are present
	return util.Message(true, "success"), true
}

// Create return new vehicle type
func (vehicleType *VehicleType) Create() (map[string]interface{}, error) {

	resp, ok := vehicleType.Validate();
	if !ok {
		return resp, nil
	}
	if err := DB.Create(vehicleType).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

// Update return upated vehicle type
func (vehicleType *VehicleType) Update(id string) (*VehicleType, error) {
	existingVehicleType, err := Find(id)
	if err != nil {
		return nil, err
	}

	if err := DB.Model(&existingVehicleType).Updates(vehicleType).Error; err != nil {
		return nil, err
	}
	return vehicleType, nil
}

// Delete delete a vehicle type
func (vehicleType *VehicleType) Delete(id string) (*VehicleType, error) {

	existingVehicleType, err := Find(id)

	if err != nil {
		return nil, err
	}
	if err := DB.Delete(existingVehicleType).Error; err != nil {
		return nil, err
	}
	return vehicleType, nil
}

// FindAll returns all vehicle type
func FindAll() ([]VehicleType, error) {
	var vtypes []VehicleType
	if err := DB.Find(&vtypes).Error; err != nil {
		return nil, err
	}
	return vtypes, nil
}

// Find returns single vehicle type
func Find(id string) (*VehicleType, error) {
	vtype := &VehicleType{}
	if err := DB.Where("id = ?", id).First(vtype).Error; err != nil {
		return nil, err
	}
	return vtype, nil
}
