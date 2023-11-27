package controllers

import (
	"math/rand"
	"parkingLot/models"
	"sync"
	"time"
)

// VehicleController struct controls the vehicles.
type VehicleController struct {
	vehicles []*models.Vehicle
}

// NewVehicleController creates a new vehicle controller.
func NewVehicleController() *VehicleController {
	return &VehicleController{
		vehicles: make([]*models.Vehicle, 0),
	}
}

// CreateVehicles creates the specified number of vehicles.
func (vc *VehicleController) CreateVehicles(numVehicles int) {
	for i := 0; i < numVehicles; i++ {
		vc.vehicles = append(vc.vehicles, models.NewVehicle(i))
	}
}

// StartSimulation starts the simulation of vehicles entering and leaving the parking lot.
// VehicleController.go
func (vc *VehicleController) StartSimulation(p *models.ParkingLot) {
	var wg sync.WaitGroup
	for _, v := range vc.vehicles {
		wg.Add(1)
		go func(v *models.Vehicle) {
			defer wg.Done()
			for {
				v.Enter(p)
				time.AfterFunc(time.Duration(rand.ExpFloat64()*1500)*time.Millisecond, func() {
					v.Leave(p)
				})
			}
		}(v)
	}
	wg.Wait()
}
