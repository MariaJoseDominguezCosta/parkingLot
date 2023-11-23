package main

import (
	"log"
	"parkingLot/controllers"
	"parkingLot/models"
	"parkingLot/scenes"
	"parkingLot/views"
)

func main() {
	// Create a new parking lot with a capacity of 20
	parkingLot := models.NewParkingLot(20)
	if parkingLot == nil {
		log.Fatal("Failed to create parking lot")
	}

	parkingController := controllers.NewParkingController(parkingLot)
	parkingView := views.NewParkingView(parkingController)

	// Start the simulation
	if err := scenes.NewParkingScene(parkingView.Window).Start(); err != nil {
		log.Fatal("Failed to start simulation:", err)
	}

	// Show the parking lot
	if err := parkingView.Show(); err != nil {
		log.Fatal("Failed to show parking lot:", err)
	}
}
