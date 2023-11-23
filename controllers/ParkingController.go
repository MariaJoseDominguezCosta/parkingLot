// controllers/ParkingController.go
package controllers

import (
	"parkingLot/models"
)

type ParkingController struct {
	ParkingLot *models.ParkingLot
}

func NewParkingController(p *models.ParkingLot) *ParkingController {
	return &ParkingController{
		ParkingLot: p,
	}
}

func (pc *ParkingController) Leave(v *models.Vehicle) {
	for i := range pc.ParkingLot.Spaces {
		if pc.ParkingLot.Spaces[i] == v {
			pc.ParkingLot.Spaces[i].Parked = true
			v.Parked = false
			break
		}
	}
}
