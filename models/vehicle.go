package models

import (
	"errors"
	"math/rand"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type VehicleId int
type Vehicle struct {
	Id              VehicleId
	Image           *canvas.Image
	ParkingDuration time.Duration
	Space           int
	Parked          bool
}

func NewVehicle(id VehicleId, imagePath string) (*Vehicle, error) {

	image := canvas.NewImageFromURI(storage.NewFileURI(imagePath))
	parkingDuration := time.Duration(10+rand.Intn(20)) * time.Second
	if id < 0 {
		return nil, errors.New("id cannot be negative")
	}
	return &Vehicle{
		Id:              id,
		ParkingDuration: parkingDuration,
		Image:           image,
		Space:           0,
		Parked:          false,
	}, nil
}
