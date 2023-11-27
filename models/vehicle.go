package models

import (
	"math/rand"
	"time"
)

// Vehicle struct represents a vehicle in the parking lot.
type Vehicle struct {
	ID int
}

// NewVehicle creates a new vehicle with the given id.
func NewVehicle(id int) *Vehicle {
	return &Vehicle{
		ID: id,
	}
}

// Enter attempts to enter the parking lot. It will block if there is no available space.
func (v *Vehicle) Enter(p *ParkingLot) {
	p.Semaphore <- v.ID                                          // Block if there is no available space.
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Stay for a random amount of time.
}

// Leave leaves the parking lot, freeing up a space.
func (v *Vehicle) Leave(p *ParkingLot) {
	<-p.Semaphore // Release the parking space.
}
