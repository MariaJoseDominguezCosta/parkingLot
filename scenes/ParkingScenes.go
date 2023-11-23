package scenes

import (
	"image/color"
	"parkingLot/controllers"
	"parkingLot/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type ParkingScene struct {
	Window fyne.Window
}

func NewParkingScene(window fyne.Window) *ParkingScene {
	return &ParkingScene{
		Window: window,
	}
}


// Show displays the parking scene.
func (s *ParkingScene) Show() error {
	rect := &canvas.Rectangle{
		StrokeWidth: 2,
		StrokeColor: color.NRGBA{R: 255, G: 255, B: 255, A: 255},
	}
	rect.Resize(fyne.NewSize(200, 300))
	rect.Move(fyne.NewPos(260, 10))

	hatch := &canvas.Rectangle{}
	hatch.Resize(fyne.NewSize(10, 160))
	hatch.Move(fyne.NewPos(195, 300))

	carsC := container.NewStack(rect, hatch)
	s.Window.SetContent(carsC)

	return nil
}

func CreateVehicle(id int) (*controllers.VehicleController, error) {
	vehicle, err := models.NewVehicle(models.VehicleId(id), "assets/car.png")
	if err != nil {
		return nil, err
	}
	return controllers.NewVehicleController(vehicle), nil
}

func (s *ParkingScene) Start() error{
	p := controllers.NewParkingController(models.NewParkingLot(20))
	poisson := models.NewPoisson()

	for i := 0; i < 100; i++ {
		vehicleController, err := CreateVehicle(i)
		if err != nil {
			// handle the error
			continue
		}

		vehicleController.Drive(p.ParkingLot)

		waitTime := poisson.Get(float64(2))
		time.Sleep(time.Duration(waitTime) * time.Second)
	}
	return nil
}
			