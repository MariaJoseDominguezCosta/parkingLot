package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fmt"
	"parkingLot/models"
	"time"
	"image/color"
	"sync"
)

type ParkingScene struct {
	Window fyne.Window

	ParkingLot *models.ParkingLot
}

func NewParkingScene(window fyne.Window, parkingLot *models.ParkingLot) *ParkingScene {
	return &ParkingScene{
		Window:     window,
		ParkingLot: parkingLot,
	}
}

var carsC = container.NewWithoutLayout()

func (s *ParkingScene) Show() {
	rect := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 0})
	rect.StrokeWidth = 2
	rect.StrokeColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	rect.Resize(fyne.NewSize(200, 300))
	rect.Move(fyne.NewPos(260,10))

	hatch := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 0})
	hatch.Resize(fyne.NewSize(10, 160))
	hatch.Move(fyne.NewPos(195,300))

	carsC.Add(rect)
	carsC.Add(hatch)
	s.Window.SetContent(carsC)
}

func (s *ParkingScene) Start() {
	p := models.NewParkingLot(make(chan int, 5), &sync.Mutex{})
	poisson := models.NewPoisson()

	var wg sync.WaitGroup

	for i := 0; i<100; i++ {
		wg.Add(1)

		go func(id int) {
			vehicle := models.NewVehicle(id)
			vehicleI := vehicle.GetVehicleImage()
			vehicleI.Resize(fyne.NewSize(50, 30))
			vehicleI.Move(fyne.NewPos(-20, 310))
			carsC.Add(vehicleI)
			carsC.Refresh()

			vehicle.Parking(p, carsC, &wg)
		}(i)
		var randPoisson = poisson.Get(float64(2))
		time.Sleep(time.Second * time.Duration(randPoisson))
	}

	wg.Wait()
	fmt.Println("Parking lot is full")
}