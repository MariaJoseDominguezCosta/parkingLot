package controllers

import (
	"parkingLot/models"
)

func OnVehicleCreated() {
	v := models.NewVehicle()

	go v.EnterParking()

}

func OnVehicleEntered() {

	models.UpdateGUI()

}
