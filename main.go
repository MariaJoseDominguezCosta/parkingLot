package main

import (
	"log"
	"parkingLot/controllers"
	"parkingLot/views"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	carInImage  *ebiten.Image
	carOutImage *ebiten.Image
)

func init() {
	var err error

	// Load the car in image file
	carInImage, _, err = ebitenutil.NewImageFromFile("assets/car.png")
	if err != nil {
		log.Fatalf("reading car in image file: %v", err)
	}

	// Load the car out image file
	carOutImage, _, err = ebitenutil.NewImageFromFile("assets/car_out.png")
	if err != nil {
		log.Fatalf("reading car out image file: %v", err)
	}
}

func main() {
	// Create a new parking lot controller with a capacity of 20.
	plc := controllers.NewParkingLotController(20)

	// Create a new vehicle controller and create 100 vehicles.
	vc := controllers.NewVehicleController()
	vc.CreateVehicles(100)

	// Start the simulation of vehicles entering and leaving the parking lot.
	vc.StartSimulation(plc.GetParkingLot())

	// Create a new parking lot view.
	plv := views.NewParkingLotView(plc)

	// Start the Ebiten game loop.
	ebiten.RunGame(plv)
}
