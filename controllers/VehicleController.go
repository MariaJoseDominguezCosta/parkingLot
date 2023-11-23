package controllers

import (
	"math/rand"
	"parkingLot/models"
	"time"
)

type VehicleController struct {
	Vehicle *models.Vehicle
	stopCh  chan struct{} // Channel to stop the goroutine
}

func NewVehicleController(v *models.Vehicle) *VehicleController {
	return &VehicleController{
		Vehicle: v,
		stopCh:  make(chan struct{}),
	}
}

func (vc *VehicleController) Drive(p *models.ParkingLot) {
	go func() {
		for {
			// Check if stop signal is received
			select {
			case <-vc.stopCh:
				return
			default:
			}

			// Simulate the arrival of vehicles according to a Poisson distribution
			waitTime := rand.ExpFloat64() * 10
			time.Sleep(time.Duration(waitTime) * time.Millisecond)

			// Try to park the vehicle in the first available space
			for i, space := range p.Spaces {
				if space == nil {
					// Park the vehicle
					p.Spaces[i] = vc.Vehicle
					vc.Vehicle.Parked = true
					parkingView.Update() // Update the view

					// Simulate the stay of the vehicle.
					parkingDuration := time.Duration(rand.Intn(20)) * time.Second
					time.Sleep(parkingDuration)

					// The vehicle leaves the parking lot
					p.Spaces[i] = nil
					vc.Vehicle.Parked = false
					parkingView.Update() // Update the view
					break
				}
			}
		}
	}()
}

func (vc *VehicleController) Stop() {
	// Send stop signal to the goroutine
	vc.stopCh <- struct{}{}
}
