package controllers

import (
	"parkingLot/models"
)

// ParkingLotController struct controls the parking lot.
type ParkingLotController struct {
	parkingLot *models.ParkingLot
}

// NewParkingLotController creates a new parking lot controller.
func NewParkingLotController(capacity int) *ParkingLotController {
	return &ParkingLotController{
		parkingLot: models.NewParkingLot(capacity),
	}
}

// GetParkingLot returns the parking lot.
func (plc *ParkingLotController) GetParkingLot() *models.ParkingLot {
	return plc.parkingLot
}
